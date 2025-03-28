package templates

import (
    "fmt"
    "os"
    "path/filepath"
    "text/template"
)

// TemplateData holds the common data for templates
type TemplateData struct {
    ProjectName    string
    PackageName    string
    Description    string
    PythonVersion  string
    // Add more common fields as needed
}

// RenderTemplate renders a template with given data to the specified path
func RenderTemplate(tmplPath, destPath string, data interface{}) error {
    // Ensure the directory exists
    dir := filepath.Dir(destPath)
    if err := os.MkdirAll(dir, 0755); err != nil {
        return fmt.Errorf("failed to create directory %s: %w", dir, err)
    }

    // Read template content
    tmpl, err := template.ParseFiles(tmplPath)
    if err != nil {
        return fmt.Errorf("failed to parse template %s: %w", tmplPath, err)
    }

    // Create destination file
    file, err := os.Create(destPath)
    if err != nil {
        return fmt.Errorf("failed to create file %s: %w", destPath, err)
    }
    defer file.Close()

    // Execute template
    if err := tmpl.Execute(file, data); err != nil {
        return fmt.Errorf("failed to execute template %s: %w", tmplPath, err)
    }

    return nil
}