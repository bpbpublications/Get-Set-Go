import (
	"log/slog"
	"os"
)
  
jsonHandler := slog.NewJSONHandler(os.Stdout, nil) // Send logs to standard output
jsonLogger := slog.New(jsonHandler)

jsonLogger.Info("Structured log: Application started")