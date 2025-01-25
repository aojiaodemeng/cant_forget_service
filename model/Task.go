package model

import (
	"cant_forget/utils"
	"cant_forget/utils/status_code"
	"github.com/mehanizm/airtable"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Id      string `gorm:"not null;unique;primary_key;comment:ID;size:90"`
	N       int
	Efactor float32
}

type CreateTaskReqParams struct {
	N        int
	Efactor  float32
	Posts    []string
	Users    []string
	Interval string
}

type TodayTodo struct {
	ColumnId       string
	ColumnName     string
	Count          int
	ColumnColorIdx float64
	FirstTwoPosts  []string
}

// 新增任务
func CreateTask(data *CreateTaskReqParams) int {
	recordsToSend := &airtable.Records{
		Records: []*airtable.Record{
			{
				Fields: map[string]any{
					"n":        data.N,
					"efactor":  data.Efactor,
					"Posts":    data.Posts,
					"Users":    data.Users,
					"interval": data.Interval,
				},
			},
		},
	}
	table := airtableClient.GetTable(utils.AirtableDBId, "Task")
	_, err := table.AddRecords(recordsToSend)
	if err != nil {
		print(err.Error())
		return status_code.ERROR // 500
	}
	return status_code.SUCCESS
}

// 查询列表
func GetTaskList() *airtable.Records {
	table := airtableClient.GetTable(utils.AirtableDBId, "Task")
	//filterFormula := fmt.Sprintf("FIND('%s', {columnsStr}) > 0", columnUuid)
	records, err := table.GetRecords().Do()

	if err != nil {
		print(err.Error())
		return nil
	}
	return records
}

// 查询今日任务列表（卡片）
func GetTodayTodoTaskList() map[string]interface{} {
	table := airtableClient.GetTable(utils.AirtableDBId, "Task")
	records, err := table.GetRecords().FromView("grouped_and_sorted").WithFilterFormula("AND({isToday}=1)").Do()
	var maps = make(map[string]any)
	// 进行分组
	for _, recordV := range records.Records {
		columnId := recordV.Fields["columns"].([]any)[0].(string)
		postTitle := recordV.Fields["postTitle"].([]any)[0].(string)
		columnName := recordV.Fields["columnTitle"].([]any)[0].(string)
		columnColorIdx := recordV.Fields["columnColorIdx"].([]any)[0]
		curColumn, exists := maps[columnId]

		if exists {
			count, _ := curColumn.(map[string]any)["count"].(int)
			firstTwoPosts, _ := curColumn.(map[string]any)["firstTwoPosts"]
			if len(firstTwoPosts.([]string)) < 2 {
				firstTwoPosts = append(firstTwoPosts.([]string), postTitle)
			}
			maps[columnId] = map[string]any{
				"columnId":       columnId,
				"columnName":     columnName,
				"count":          count + 1,
				"columnColorIdx": columnColorIdx,
				"firstTwoPosts":  firstTwoPosts,
			}

		} else {
			var arr []string
			arr = append(arr, postTitle)
			maps[columnId] = map[string]any{
				"columnId":       columnId,
				"columnName":     columnName,
				"count":          1,
				"columnColorIdx": columnColorIdx,
				"firstTwoPosts":  arr,
			}
		}
	}

	if err != nil {
		print(err.Error())
		return nil
	}
	return maps
}

// 查询今日任务列表
func GetTodoList() []*airtable.Record {
	table := airtableClient.GetTable(utils.AirtableDBId, "Task")
	records, err := table.GetRecords().WithFilterFormula("AND({isToday}=1)").Do()
	if err != nil {
		print(err.Error())
		return nil
	}
	return records.Records
}

func AddColumnToTask(columnId string) int {
	columnTable := airtableClient.GetTable(utils.AirtableDBId, "Column")
	taskTable := airtableClient.GetTable(utils.AirtableDBId, "Task")
	postTable := airtableClient.GetTable(utils.AirtableDBId, "Post")
	record, err := columnTable.GetRecord(columnId)
	if err != nil {
		print(err.Error())
		return status_code.ERROR
	}
	fields := record.Fields
	var records []*airtable.Record
	for _, postId := range fields["Posts"].([]interface{}) {
		var users []string
		users = append(users, "rechGT5IRtq9Hdgf4")
		var posts []string
		posts = append(posts, postId.(string))
		postRecord, _ := postTable.GetRecord(postId.(string))
		postFields := postRecord.Fields

		curRecord := &airtable.Record{
			Fields: map[string]any{
				"n":        0,
				"efactor":  postFields["efactor"],
				"Posts":    posts,
				"Users":    users,
				"interval": time.Now(),
			},
		}
		records = append(records, curRecord)
	}

	recordsToSend := &airtable.Records{
		Records: records,
	}

	_, err2 := taskTable.AddRecords(recordsToSend)
	if err2 != nil {
		print(err.Error())
		return status_code.ERROR // 500
	}
	return status_code.SUCCESS
}
