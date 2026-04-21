package mapper

import (
	"time"

	"github.com/google/uuid"
	
	"cli-todo-app-golang/internal/interface/dto"
	"cli-todo-app-golang/internal/domain/entity"
)

func ToUserStorageFromEntity(u entity.User) dto.UserStorage {
	return dto.UserStorage{
		UserId:    u.UserId,
		Name:      u.Name,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	}
}

func ToEntityFromUserStorage(u dto.UserStorage) entity.User {
	return entity.User{
		UserId:    u.UserId,
		Name:      u.Name,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	}
}

func ToUserLogStorageFromEntity(u entity.UserLog) dto.UserLogStorage {
	return dto.UserLogStorage{
		LogId:  u.LogId,
		UserId: u.UserId,
		Action: u.Action,
		Time:   u.Time,
		
	}
}

func ToEntityFromUserLogStorage(u dto.UserLogStorage) entity.UserLog {
	return entity.UserLog{
		LogId:  u.LogId,
		UserId: u.UserId,
		Action: u.Action,
		Time:   u.Time,
		
	}
}

func ToSessionStorageFromEntity(u entity.Session) dto.SessionStorage {
	return dto.SessionStorage{
		UserId:    u.UserId,
		Name:      u.Name,
		Username:  u.Username,
		Email:     u.Email,
		LoggedIn:  u.LoggedIn,
		LoginTime: u.LoginTime,
		
	}
}

func ToEntityFromSessionStorage(u dto.SessionStorage) entity.Session {
	return entity.Session{
		UserId:    u.UserId,
		Name:      u.Name,
		Username:  u.Username,
		Email:     u.Email,
		LoggedIn:  u.LoggedIn,
		LoginTime: u.LoginTime,
		
	}
}

func ToTodoStorageFromEntity(userId string, todos []entity.Todo) dto.TodoStorage {
    items := make([]dto.TodoItem, len(todos))
    for i, u := range todos {
        items[i] = dto.TodoItem{
            ID:        u.ID,
            Task:      u.Task,
            DueDate:   u.DueDate,
            Priority:  u.Priority,
            Progress:  u.Progress,
            CreatedAt: u.CreatedAt,
        }
    }
    return dto.TodoStorage{
        UserId: userId,
        Todos:  items,
    }
}

func ToEntityFromTodoStorage(storage dto.TodoStorage) (string, []entity.Todo) {
    todos := make([]entity.Todo, len(storage.Todos))
    for i, item := range storage.Todos {
        todos[i] = entity.Todo{
            ID:        item.ID,
            UserId:    storage.UserId,
            Task:      item.Task,
            DueDate:   item.DueDate,
            Priority:  item.Priority,
            Progress:  item.Progress,
            CreatedAt: item.CreatedAt,
        }
    }
    return storage.UserId, todos
}

func ToEntityFromAddTodoRequest(userId string, req dto.AddTodoRequest) entity.Todo {
    dueDate, _ := time.Parse("2006-01-02", req.DueDate)

    return entity.Todo{
        ID:        uuid.New().String(),
        UserId:    userId,
        Task:      req.Task,
        Priority:  req.Priority,
        Progress:  "pending",
        DueDate:   dueDate,
        CreatedAt: time.Now(),
    }
}

func ToEntityFromRegisterRequest(password string, req dto.RegisterRequest) entity.User {
	return entity.User{
		UserId:    uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Username:  req.Username,
		Password:  password,
		CreatedAt: time.Now(),
	}
} 

func ToEntityFromSession(userId, name, username, email string, loginTime time.Time) entity.Session {
    return entity.Session{
        UserId:    userId,
        Name:      name,
        Username:  username,
        Email:     email,
        LoggedIn:  true,
        LoginTime: loginTime,
    }
}

func ToEntityFromNewLogUser(userId, action string) entity.UserLog {
	return entity.UserLog{
		LogId:  uuid.New().String(),
		UserId: userId,
		Action: action,
		Time:   time.Now(),
	} 
}