customSlogHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelDebug,
})
customSLogger := slog.New(customSlogHandler)
customSLogger.Debug("Debugging information", "feature", "new-logging", "status", "active")
customSLogger.Warn("This is a warning message", "module", "logging", "severity", "high")
customSLogger.Error("Error encountered", "error", "file-not-found", "retry", true)
