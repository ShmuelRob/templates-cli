package templates

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    
    "github.com/urfave/cli/v2"
)

// ETLTemplateData holds data for ETL template generation
type ETLTemplateData struct {
    TemplateData
    ExtractMethod   string
    TransformMethod string
    LoadDestination string
    Dependencies    []string
}

// GenerateETLTemplate generates a Python ETL project template
func GenerateETLTemplate(c *cli.Context) error {
    projectName := c.String("name")
    extractMethod := c.String("extract")
    transformMethod := c.String("transform")
    loadDestination := c.String("load")
    createVenv := c.Bool("venv")
    
    // Convert project name to package name (lowercase, replace hyphens with underscores)
    packageName := strings.ReplaceAll(strings.ToLower(projectName), "-", "_")
    
    // Validate inputs
    if err := validateETLInputs(extractMethod, transformMethod, loadDestination); err != nil {
        return err
    }
    
    // Create project directory
    if err := os.MkdirAll(projectName, 0755); err != nil {
        return fmt.Errorf("failed to create project directory: %w", err)
    }
    
    // Determine dependencies based on components
    dependencies := determineDependencies(extractMethod, transformMethod, loadDestination)
    
    // Prepare template data
    data := ETLTemplateData{
        TemplateData: TemplateData{
            ProjectName:   projectName,
            PackageName:   packageName,
            Description:   "A Python ETL project for data processing",
            PythonVersion: ">=3.8",
        },
        ExtractMethod:   extractMethod,
        TransformMethod: transformMethod,
        LoadDestination: loadDestination,
        Dependencies:    dependencies,
    }
    
    // Get all template files and directories
    if err := generateProjectFiles(projectName, data); err != nil {
        return err
    }
    
    // Initialize virtual environment if requested
    if createVenv {
        if err := initializeVirtualEnv(projectName); err != nil {
            return err
        }
    }
    
    fmt.Printf("Python ETL project template generated successfully in %s\n", projectName)
    return nil
}

// validateETLInputs validates the ETL input parameters
func validateETLInputs(extract, transform, load string) error {
    validExtractMethods := map[string]bool{"file": true, "api": true, "database": true}
    validTransformMethods := map[string]bool{"basic": true, "advanced": true}
    validLoadDestinations := map[string]bool{"file": true, "database": true, "api": true}
    
    if !validExtractMethods[extract] {
        return fmt.Errorf("invalid extract method: %s. Valid options are: file, api, database", extract)
    }
    
    if !validTransformMethods[transform] {
        return fmt.Errorf("invalid transform method: %s. Valid options are: basic, advanced", transform)
    }
    
    if !validLoadDestinations[load] {
        return fmt.Errorf("invalid load destination: %s. Valid options are: file, database, api", load)
    }
    
    return nil
}

// determineDependencies returns a list of package dependencies based on selected methods
func determineDependencies(extract, transform, load string) []string {
    dependencies := []string{"pytest", "python-dotenv"}
    
    // Extract dependencies
    switch extract {
    case "file":
        dependencies = append(dependencies, "pandas")
    case "api":
        dependencies = append(dependencies, "requests")
    case "database":
        dependencies = append(dependencies, "sqlalchemy", "psycopg2-binary")
    }
    
    // Transform dependencies
    switch transform {
    case "basic":
        dependencies = append(dependencies, "pandas")
    case "advanced":
        dependencies = append(dependencies, "pandas", "numpy", "scikit-learn")
    }
    
    // Load dependencies
    switch load {
    case "file":
        // pandas already added
    case "database":
        if !contains(dependencies, "sqlalchemy") {
            dependencies = append(dependencies, "sqlalchemy", "psycopg2-binary")
        }
    case "api":
        if !contains(dependencies, "requests") {
            dependencies = append(dependencies, "requests")
        }
    }
    
    return dependencies
}

// contains checks if a string is in a slice
func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}

// generateProjectFiles generates all project files from templates
func generateProjectFiles(projectName string, data ETLTemplateData) error {
    templateDir := "templates/etl-python"
    
    // Define the mapping of template files to destination paths
    templatesMap := map[string]string{
        "README.md.tmpl":                      filepath.Join(projectName, "README.md"),
        "requirements.txt.tmpl":               filepath.Join(projectName, "requirements.txt"),
        "setup.py.tmpl":                       filepath.Join(projectName, "setup.py"),
        ".gitignore.tmpl":                     filepath.Join(projectName, ".gitignore"),
        "src/__init__.py.tmpl":                filepath.Join(projectName, "src", "__init__.py"),
        "src/main.py.tmpl":                    filepath.Join(projectName, "src", "main.py"),
        "src/extract/__init__.py.tmpl":        filepath.Join(projectName, "src", "extract", "__init__.py"),
        "src/extract/extract.py.tmpl":         filepath.Join(projectName, "src", "extract", "extract.py"),
        "src/transform/__init__.py.tmpl":      filepath.Join(projectName, "src", "transform", "__init__.py"),
        "src/transform/transform.py.tmpl":     filepath.Join(projectName, "src", "transform", "transform.py"),
        "src/load/__init__.py.tmpl":           filepath.Join(projectName, "src", "load", "__init__.py"),
        "src/load/load.py.tmpl":               filepath.Join(projectName, "src", "load", "load.py"),
        "tests/__init__.py.tmpl":              filepath.Join(projectName, "tests", "__init__.py"),
        "tests/test_extract.py.tmpl":          filepath.Join(projectName, "tests", "test_extract.py"),
        "tests/test_transform.py.tmpl":        filepath.Join(projectName, "tests", "test_transform.py"),
        "tests/test_load.py.tmpl":             filepath.Join(projectName, "tests", "test_load.py"),
    }
    
    // Render each template
    for tmpl, dest := range templatesMap {
        if err := RenderTemplate(filepath.Join(templateDir, tmpl), dest, data); err != nil {
            return err
        }
    }
    
    return nil
}

// initializeVirtualEnv creates and initializes a Python virtual environment
func initializeVirtualEnv(projectDir string) error {
    fmt.Println("Initializing Python virtual environment...")
    
    // Create virtual environment
    cmd := exec.Command("python", "-m", "venv", filepath.Join(projectDir, "venv"))
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to create virtual environment: %w", err)
    }
    
    // Install dependencies
    var pipCmd *exec.Cmd
    if runtime.GOOS == "windows" {
        pipCmd = exec.Command(filepath.Join(projectDir, "venv", "Scripts", "pip"), "install", "-r", filepath.Join(projectDir, "requirements.txt"))
    } else {
        pipCmd = exec.Command(filepath.Join(projectDir, "venv", "bin", "pip"), "install", "-r", filepath.Join(projectDir, "requirements.txt"))
    }
    
    pipCmd.Stdout = os.Stdout
    pipCmd.Stderr = os.Stderr
    if err := pipCmd.Run(); err != nil {
        return fmt.Errorf("failed to install dependencies: %w", err)
    }
    
    return nil
}