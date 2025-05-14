customSlogHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	AddSource: true, // Include source file and line number
  })
customSLogger := slog.New(customSlogHandler)
customSLogger.Info("Log with source information")
  