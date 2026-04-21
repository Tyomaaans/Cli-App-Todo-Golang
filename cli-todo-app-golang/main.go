package main

import (
    "cli-todo-app-golang/internal/infrastructure/storage/json"
    "cli-todo-app-golang/internal/infrastructure/config"
    "cli-todo-app-golang/internal/interface/cli"

    ucauth "cli-todo-app-golang/internal/usecase/auth"
    ucuser "cli-todo-app-golang/internal/usecase/user"
    uctodo "cli-todo-app-golang/internal/usecase/todo"

    cliauth "cli-todo-app-golang/internal/interface/cli/auth"
    cliuser "cli-todo-app-golang/internal/interface/cli/user"
    clitodo "cli-todo-app-golang/internal/interface/cli/todo"

    appvalidator "cli-todo-app-golang/internal/validator"
)

func main() {
    validate := appvalidator.New()
    
    pathConfig := config.NewPathConfig("data")
    
    sessionRepo := json.NewSessionRepository(pathConfig.SessionFile())
    userRepo    := json.NewUserRepository(pathConfig.UserFile())
    userLogRepo := json.NewLogRepository(pathConfig.UserLogFile())
    todoRepo    := json.NewTodoRepository(pathConfig.TodoFile())

    loginUC          := ucauth.NewLoginUseCase(userRepo, sessionRepo, userLogRepo, validate)
    logoutUC         := ucauth.NewLogoutUseCase(sessionRepo, userLogRepo)
    registerUC       := ucauth.NewRegisterUseCase(userRepo, sessionRepo, validate)
    forgotpasswordUC := ucauth.NewForgotPasswordUseCase(userRepo, sessionRepo, userLogRepo, validate)

    updatenameUC     := ucuser.NewUpdateNameUseCase(userRepo, sessionRepo, userLogRepo, validate)
    updateemailUC    := ucuser.NewUpdateEmailUseCase(userRepo, sessionRepo, userLogRepo, validate)
    updateusernameUC := ucuser.NewUpdateUsernameUseCase(userRepo, sessionRepo, userLogRepo, validate)
    updatepasswordUC := ucuser.NewUpdatePasswordUseCase(userRepo, sessionRepo, userLogRepo, validate)
    updateuserUC     := ucuser.NewUpdateUserUseCase(userRepo, sessionRepo, userLogRepo, validate)
    deleteuserUC     := ucuser.NewDeleteUserUseCase(userRepo, sessionRepo, userLogRepo)

    addtodoUC        := uctodo.NewAddTodoUseCase(todoRepo, sessionRepo, userLogRepo, validate)
    deletetodoUC     := uctodo.NewDeleteTodoUseCase(todoRepo, sessionRepo, userLogRepo, validate)
    listtodoUC       := uctodo.NewListTodoUseCase(todoRepo, sessionRepo, userLogRepo)
    updatetodoUC     := uctodo.NewUpdateTodoUseCase(todoRepo, sessionRepo, userLogRepo, validate)

    loginHandler          := cliauth.NewLoginHandler(loginUC)
    logoutHandler         := cliauth.NewLogoutHandler(logoutUC)
    registerHandler       := cliauth.NewRegisterHandler(registerUC)
    forgotpasswordHandler := cliauth.NewForgotPasswordHandler(forgotpasswordUC)

    updatenameHandler     := cliuser.NewUpdateNameHandler(updatenameUC)
    updateemailHandler    := cliuser.NewUpdateEmailHandler(updateemailUC)
    updateusernameHandler := cliuser.NewUpdateUsernameHandler(updateusernameUC)
    updatepasswordHandler := cliuser.NewUpdatePaswordHandler(updatepasswordUC)
    updateuserHandler     := cliuser.NewUpdateUserHandler(updateuserUC)
    deleteuserHandler     := cliuser.NewDeleteUserlHandler(deleteuserUC)

    addtodoHandler        := clitodo.NewAddTodoHandler(addtodoUC)
    deletetodoHandler     := clitodo.NewDeleteTodoHandler(deletetodoUC)
    listtodoHandler       := clitodo.NewListTodoHandler(listtodoUC)
    updatetodoHandler     := clitodo.NewUpdateTodoHandler(updatetodoUC)

    app := cli.NewApp(
        loginHandler, 
        logoutHandler, 
        registerHandler, 
        forgotpasswordHandler, 

        updatenameHandler, 
        updateemailHandler,
        updateusernameHandler,
        updatepasswordHandler,
        updateuserHandler,
        deleteuserHandler,

        addtodoHandler,
        deletetodoHandler,
        listtodoHandler,
        updatetodoHandler,
    )

    app.Run()
}