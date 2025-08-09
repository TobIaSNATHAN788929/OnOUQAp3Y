// 代码生成时间: 2025-08-09 17:37:38
package main

import (
    "fmt"
    "os/exec"
    "os"
    "time"
    "context"
    "github.com/kataras/iris/v12"
)

// ProcessManager is a struct that manages processes
type ProcessManager struct {
    // Context for managing processes
    ctx context.Context
}

// NewProcessManager creates a new instance of ProcessManager
func NewProcessManager() *ProcessManager {
    return &ProcessManager{
        ctx: context.Background(),
    }
}

// StartProcess starts a new process with the given command
func (pm *ProcessManager) StartProcess(command string) (*os.Process, error) {
    // Split the command into separate parts
    args := strings.Fields(command)
    if len(args) == 0 {
        return nil, fmt.Errorf("invalid command")
    }

    // Create a new process
    process, err := os.StartProcess(args[0], args, &os.ProcAttr{})
    if err != nil {
        return nil, err
    }

    return process, nil
}

// StopProcess stops the given process
func (pm *ProcessManager) StopProcess(process *os.Process) error {
    // Send a SIGTERM signal to the process
    if err := process.Signal(os.Interrupt); err != nil {
        return err
    }

    // Wait for the process to exit
    _, err := process.Wait()
    return err
}

// StartProcessManagerRoute sets up the route for managing processes
func StartProcessManagerRoute(app *iris.Application) {
    pm := NewProcessManager()

    // Route to start a new process
    app.Post("/start", func(ctx iris.Context) {
        command := ctx.URLParam("command")
        process, err := pm.StartProcess(command)
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }

        ctx.JSON(iris.Map{
            "message": "Process started",
            "pid": process.Pid,
        })
    })

    // Route to stop a process
    app.Post("/stop", func(ctx iris.Context) {
        pid := ctx.URLParam("pid")
        process, err := os.FindProcess(pid)
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }

        if err := pm.StopProcess(process); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }

        ctx.JSON(iris.Map{
            "message": "Process stopped",
        })
    })
}

func main() {
    app := iris.New()
    StartProcessManagerRoute(app)
    app.Listen(":8080")
}
