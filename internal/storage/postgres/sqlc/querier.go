// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateHeading(ctx context.Context, arg CreateHeadingParams) error
	CreateList(ctx context.Context, arg CreateListParams) error
	CreateTag(ctx context.Context, arg CreateTagParams) error
	CreateTask(ctx context.Context, arg CreateTaskParams) error
	DeleteHeading(ctx context.Context, arg DeleteHeadingParams) error
	DeleteList(ctx context.Context, arg DeleteListParams) error
	GetArchivedTasks(ctx context.Context, arg GetArchivedTasksParams) ([]GetArchivedTasksRow, error)
	GetCompletedTasks(ctx context.Context, arg GetCompletedTasksParams) ([]GetCompletedTasksRow, error)
	GetDefaultHeadingID(ctx context.Context, arg GetDefaultHeadingIDParams) (string, error)
	GetDefaultListID(ctx context.Context, userID string) (string, error)
	GetHeadingByID(ctx context.Context, arg GetHeadingByIDParams) (GetHeadingByIDRow, error)
	GetHeadingsByListID(ctx context.Context, arg GetHeadingsByListIDParams) ([]GetHeadingsByListIDRow, error)
	GetListByID(ctx context.Context, arg GetListByIDParams) (GetListByIDRow, error)
	GetListsByUserID(ctx context.Context, userID string) ([]GetListsByUserIDRow, error)
	GetOverdueTasks(ctx context.Context, arg GetOverdueTasksParams) ([]GetOverdueTasksRow, error)
	GetTagIDByTitle(ctx context.Context, arg GetTagIDByTitleParams) (string, error)
	GetTagsByTaskID(ctx context.Context, taskID string) ([]GetTagsByTaskIDRow, error)
	GetTagsByUserID(ctx context.Context, userID string) ([]GetTagsByUserIDRow, error)
	GetTaskByID(ctx context.Context, arg GetTaskByIDParams) (GetTaskByIDRow, error)
	GetTaskStatusID(ctx context.Context, title string) (int32, error)
	GetTasksByListID(ctx context.Context, arg GetTasksByListIDParams) ([]GetTasksByListIDRow, error)
	GetTasksByUserID(ctx context.Context, arg GetTasksByUserIDParams) ([]GetTasksByUserIDRow, error)
	GetTasksForSomeday(ctx context.Context, arg GetTasksForSomedayParams) ([]GetTasksForSomedayRow, error)
	GetTasksForToday(ctx context.Context, userID string) ([]GetTasksForTodayRow, error)
	GetTasksGroupedByHeadings(ctx context.Context, arg GetTasksGroupedByHeadingsParams) ([]GetTasksGroupedByHeadingsRow, error)
	GetUpcomingTasks(ctx context.Context, arg GetUpcomingTasksParams) ([]GetUpcomingTasksRow, error)
	LinkTagToTask(ctx context.Context, arg LinkTagToTaskParams) error
	MarkTaskAsArchived(ctx context.Context, arg MarkTaskAsArchivedParams) error
	MarkTaskAsCompleted(ctx context.Context, arg MarkTaskAsCompletedParams) error
	MoveHeadingToAnotherList(ctx context.Context, arg MoveHeadingToAnotherListParams) error
	MoveTaskToAnotherList(ctx context.Context, arg MoveTaskToAnotherListParams) error
	UnlinkTagFromTask(ctx context.Context, arg UnlinkTagFromTaskParams) error
	UpdateHeading(ctx context.Context, arg UpdateHeadingParams) error
	UpdateList(ctx context.Context, arg UpdateListParams) error
	UpdateTasksListID(ctx context.Context, arg UpdateTasksListIDParams) error
}

var _ Querier = (*Queries)(nil)
