package cli

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"

    cliauth "cli-todo-app-golang/internal/interface/cli/auth"
    cliuser "cli-todo-app-golang/internal/interface/cli/user"
    clitodo "cli-todo-app-golang/internal/interface/cli/todo"
    
    cmdauth "cli-todo-app-golang/internal/interface/cmd/auth"
    cmduser "cli-todo-app-golang/internal/interface/cmd/user"
    cmdtodo "cli-todo-app-golang/internal/interface/cmd/todo"
)

type App struct {
    loginHandler          *cliauth.LoginHandler
    logoutHandler         *cliauth.LogoutHandler
    registerHandler       *cliauth.RegisterHandler
    forgotpasswordHandler *cliauth.ForgotPasswordHandler

    updatenameHandler     *cliuser.UpdateNameHandler
    updateemailHandler    *cliuser.UpdateEmailHandler
    updateusernameHandler *cliuser.UpdateUsernameHandler
    updatepasswordHandler *cliuser.UpdatePasswordHandler
    updateuserHandler     *cliuser.UpdateUserHandler
    deleteuserHandler     *cliuser.DeleteUserHandler

    addtodoHandler        *clitodo.AddTodoHandler
    deletetodoHandler     *clitodo.DeleteTodoHandler
    listtodoHandler       *clitodo.ListTodoHandler
    updatetodoHandler     *clitodo.UpdateTodoHandler
}

func NewApp(
    loginHandler          *cliauth.LoginHandler,
    logoutHandler         *cliauth.LogoutHandler,
    registerHandler       *cliauth.RegisterHandler,
    forgotpasswordHandler *cliauth.ForgotPasswordHandler,

    updatenameHandler     *cliuser.UpdateNameHandler,
    updateemailHandler    *cliuser.UpdateEmailHandler,
    updateusernameHandler *cliuser.UpdateUsernameHandler,
    updatepasswordHandler *cliuser.UpdatePasswordHandler,
    updateuserHandler     *cliuser.UpdateUserHandler,
    deleteuserHandler     *cliuser.DeleteUserHandler,

    addtodoHandler        *clitodo.AddTodoHandler,
    deletetodoHandler     *clitodo.DeleteTodoHandler,
    listtodoHandler       *clitodo.ListTodoHandler,
    updatetodoHandler     *clitodo.UpdateTodoHandler,
) *App {
    return &App{
        loginHandler:          loginHandler,
        logoutHandler:         logoutHandler,
        registerHandler:       registerHandler,
        forgotpasswordHandler: forgotpasswordHandler,

        updatenameHandler:     updatenameHandler,
        updateemailHandler:    updateemailHandler,
        updateusernameHandler: updateusernameHandler,
        updatepasswordHandler: updatepasswordHandler,
        updateuserHandler:     updateuserHandler,
        deleteuserHandler:     deleteuserHandler,

        addtodoHandler:        addtodoHandler,
        deletetodoHandler:     deletetodoHandler,
        listtodoHandler:       listtodoHandler,
        updatetodoHandler:     updatetodoHandler,
    }
}

func (a *App) Run() {
    rootCmd := &cobra.Command{
        Use: "todo-cli",
    }

    loginCmd          := cmdauth.NewLoginCommand(a.loginHandler.Login)
    logoutCmd         := cmdauth.NewLogoutCommand(a.logoutHandler.Logout)
    registerCmd       := cmdauth.NewRegisterCommand(a.registerHandler.Register)
    forgotpasswordCmd := cmdauth.NewForgotPasswordCommand(a.forgotpasswordHandler.ForgotPassword)

    userCmd := &cobra.Command{
        Use:   "user",
        Short: "Manage user profile",
    }

    updateNameCmd     := cmduser.NewUpdateNameCommand(a.updatenameHandler.UpdateName)
    updateEmailCmd    := cmduser.NewUpdateEmailCommand(a.updateemailHandler.UpdateEmail)
    updateUsernameCmd := cmduser.NewUpdateUsernameCommand(a.updateusernameHandler.UpdateUsername)
    updatePasswordCmd := cmduser.NewUpdatePasswordCommand(a.updatepasswordHandler.UpdatePassword)
    updateUserCmd     := cmduser.NewUpdateUserCommand(a.updateuserHandler.UpdateUser)
    deleteUserCmd     := cmduser.NewDeleteUserCommand(a.deleteuserHandler.DeleteUser)

    todoCmd := &cobra.Command{
        Use:   "todo",
        Short: "Manage todo",
    }

    addTodoCmd    := cmdtodo.NewAddTodoCommand(a.addtodoHandler.AddTodo)
    deleteTodoCmd := cmdtodo.NewDeleteTodoCommand(a.deletetodoHandler.DeleteTodo)
    listTodoCmd   := cmdtodo.NewListTodoCommand(a.listtodoHandler.ListTodo)
    updateTodoCmd := cmdtodo.NewUpdateTodoCommand(a.updatetodoHandler.UpdateTodo)

    rootCmd.AddCommand(loginCmd)
    rootCmd.AddCommand(logoutCmd)
    rootCmd.AddCommand(registerCmd)
    rootCmd.AddCommand(forgotpasswordCmd)
    
    rootCmd.AddCommand(userCmd)
    userCmd.AddCommand(updateNameCmd)
    userCmd.AddCommand(updateEmailCmd)
    userCmd.AddCommand(updateUsernameCmd)
    userCmd.AddCommand(updatePasswordCmd)
    userCmd.AddCommand(updateUserCmd)
    userCmd.AddCommand(deleteUserCmd)

    rootCmd.AddCommand(todoCmd)
    todoCmd.AddCommand(addTodoCmd)
    todoCmd.AddCommand(deleteTodoCmd)
    todoCmd.AddCommand(listTodoCmd)
    todoCmd.AddCommand(updateTodoCmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}