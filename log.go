package fiberEvents

import "github.com/gofiber/fiber/v2"

type Log struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func eventLogAppend(c *fiber.Ctx, level string, message string) {
	event := c.Locals("event").(*Event)

	log := &Log{
		Level:   level,
		Message: message,
	}

	event.Tracelog = append(event.Tracelog, log)
}

func Emergency(c *fiber.Ctx, message string) {
	eventLogAppend(c, "emergency", message)
}

func Alert(c *fiber.Ctx, message string) {
	eventLogAppend(c, "alert", message)
}

func Critical(c *fiber.Ctx, message string) {
	eventLogAppend(c, "critical", message)
}

func Error(c *fiber.Ctx, message string) {
	eventLogAppend(c, "error", message)
}

func Warning(c *fiber.Ctx, message string) {
	eventLogAppend(c, "warning", message)
}

func Notice(c *fiber.Ctx, message string) {
	eventLogAppend(c, "notice", message)
}

func Info(c *fiber.Ctx, message string) {
	eventLogAppend(c, "info", message)
}

func Trace(c *fiber.Ctx, message string) {
	eventLogAppend(c, "trace", message)
}

func Debug(c *fiber.Ctx, message string) {
	eventLogAppend(c, "debug", message)
}
