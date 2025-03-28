package utils

import (
    "os"
    "runtime"
)

// GetTemplatesDir returns the absolute path to the templates directory
func GetTemplatesDir() string {
    // For development, use the local templates dir
    if _, err := os.Stat("templates"); err == nil {
        return "templates"
    }
    
    // For installed application, determine templates directory based on OS
    switch runtime.GOOS {
    case "windows":
        // Windows typically uses %APPDATA%
        return filepath.Join(os.Getenv("APPDATA"), "pytgen", "templates")
    default:
        // Linux/macOS typically uses ~/.config
        homeDir, _ := os.UserHomeDir()
        return filepath.Join(homeDir, ".config", "pytgen", "templates")
    }
}