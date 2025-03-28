package templates

import (
    "fmt"
    "os"
	"strings"
    "github.com/AlecAivazis/survey/v2"
    "github.com/urfave/cli/v2"
)

// ProjectType represents different project template types
type ProjectType string

const (
    ETLProject ProjectType = "ETL (Extract, Transform, Load)"
    // Add more project types here as you expand
)

// InteractiveGenerator launches an interactive project generator
func InteractiveGenerator(c *cli.Context) error {
    fmt.Println("🚀 Welcome to the Python Project Template Generator!")
    fmt.Println("This wizard will guide you through creating a new project.")
    
    // Step 1: Choose project type
    projectType, err := promptProjectType()
    if err != nil {
        return err
    }
    
    // Step 2: Get project details based on type
    switch projectType {
    case ETLProject:
        return promptETLProjectDetails()
    default:
        return fmt.Errorf("unsupported project type: %s", projectType)
    }
}

// promptProjectType asks the user to select a project type
func promptProjectType() (ProjectType, error) {
    var projectType ProjectType
    
    prompt := &survey.Select{
        Message: "What type of project do you want to create?",
        Options: []string{
            string(ETLProject),
            // Add more project types here
        },
        Description: func(value string, index int) string {
            switch value {
            case string(ETLProject):
                return "A data pipeline for extracting, transforming, and loading data"
            default:
                return ""
            }
        },
    }
    
    err := survey.AskOne(prompt, &projectType)
    return projectType, err
}

// promptETLProjectDetails collects details for an ETL project
// func promptETLProjectDetails() error {
//     // Questions for ETL project
//     questions := []*survey.Question{
//         {
//             Name: "projectName",
//             Prompt: &survey.Input{
//                 Message: "Project name:",
//                 Default: "python-etl-project",
//                 Help:    "The name of your project directory and package",
//             },
//             Validate: survey.Required,
//         },
//         {
//             Name: "extractMethod",
//             Prompt: &survey.Select{
//                 Message: "Select extraction method:",
//                 Options: []string{"file", "api", "database"},
//                 Default: "file",
//                 Description: func(value string, index int) string {
//                     switch value {
//                     case "file":
//                         return "Extract data from CSV, Excel, or other files"
//                     case "api":
//                         return "Extract data from REST APIs"
//                     case "database":
//                         return "Extract data from SQL databases"
//                     default:
//                         return ""
//                     }
//                 },
//             },
//         },
//         {
//             Name: "transformMethod",
//             Prompt: &survey.Select{
//                 Message: "Select transformation method:",
//                 Options: []string{"basic", "advanced"},
//                 Default: "basic",
//                 Description: func(value string, index int) string {
//                     switch value {
//                     case "basic":
//                         return "Simple data cleaning and formatting"
//                     case "advanced":
//                         return "Advanced processing including feature engineering, scaling, etc."
//                     default:
//                         return ""
//                     }
//                 },
//             },
//         },
//         {
//             Name: "loadDestination",
//             Prompt: &survey.Select{
//                 Message: "Select load destination:",
//                 Options: []string{"file", "database", "api"},
//                 Default: "file",
//                 Description: func(value string, index int) string {
//                     switch value {
//                     case "file":
//                         return "Load data to CSV, Excel, or other files"
//                     case "database":
//                         return "Load data to SQL databases"
//                     case "api":
//                         return "Load data to REST APIs"
//                     default:
//                         return ""
//                     }
//                 },
//             },
//         },
//         {
//             Name: "createVenv",
//             Prompt: &survey.Confirm{
//                 Message: "Initialize a virtual environment?",
//                 Default: false,
//                 Help:    "Creates a Python virtual environment and installs dependencies",
//             },
//         },
//     }
    
//     // Store answers
//     answers := struct {
//         ProjectName     string
//         ExtractMethod   string
//         TransformMethod string
//         LoadDestination string
//         CreateVenv      bool
//     }{}
    
//     // Ask the questions
//     err := survey.Ask(questions, &answers)
//     if err != nil {
//         return err
//     }

// 	// Add to the promptETLProjectDetails function

// 	// Ask for specific details based on extract method
// 	switch answers.ExtractMethod {
// 	case "file":
// 		var fileType string
// 		filePrompt := &survey.Select{
// 			Message: "What type of files will you extract from?",
// 			Options: []string{"CSV", "Excel", "JSON", "Parquet", "Other"},
// 			Default: "CSV",
// 		}
// 		survey.AskOne(filePrompt, &fileType)
		
// 		// You could store this in the template data or use additional templates
		
// 	case "database":
// 		var dbType string
// 		dbPrompt := &survey.Select{
// 			Message: "What database will you extract from?",
// 			Options: []string{"PostgreSQL", "MySQL", "SQLite", "Oracle", "SQL Server", "Other"},
// 			Default: "PostgreSQL",
// 		}
// 		survey.AskOne(dbPrompt, &dbType)
		
// 		// You could generate database-specific connection code
// 	}

// 	// Add to the promptETLProjectDetails function

// 	additionalDeps := []string{}
// 	depsPrompt := &survey.MultiSelect{
// 		Message: "Select additional dependencies to include:",
// 		Options: []string{
// 			"matplotlib - Plotting library",
// 			"seaborn - Statistical visualization",
// 			"pytest-cov - Test coverage",
// 			"black - Code formatter",
// 			"mypy - Type checking",
// 			"pydantic - Data validation",
// 		},
// 	}
// 	survey.AskOne(depsPrompt, &additionalDeps)

// 	// Process the selected dependencies
// 	for _, dep := range additionalDeps {
// 		// Extract just the package name before the dash
// 		packageName := strings.Split(dep, " - ")[0]
// 		etlData.Dependencies = append(etlData.Dependencies, packageName)
// 	}
		
//     // Show summary
//     fmt.Println("\n📋 Project Summary:")
//     fmt.Printf("  • Name: %s\n", answers.ProjectName)
//     fmt.Printf("  • Extract: %s\n", answers.ExtractMethod)
//     fmt.Printf("  • Transform: %s\n", answers.TransformMethod)
//     fmt.Printf("  • Load: %s\n", answers.LoadDestination)
//     fmt.Printf("  • Virtual Environment: %v\n", answers.CreateVenv)
    
//     // Confirm generation
//     proceed := false
//     prompt := &survey.Confirm{
//         Message: "Generate this project?",
//         Default: true,
//     }
//     survey.AskOne(prompt, &proceed)
    
//     if !proceed {
//         fmt.Println("Project generation cancelled.")
//         return nil
//     }
    
//     // Generate the template
//     fmt.Println("\n🔨 Generating project...")
    
//     // Use existing template generator
//     etlData := ETLTemplateData{
//         TemplateData: TemplateData{
//             ProjectName:  answers.ProjectName,
//             PackageName:  answers.ProjectName,
//             Description:  "A Python ETL project for data processing",
//             PythonVersion: ">=3.8",
//         },
//         ExtractMethod:   answers.ExtractMethod,
//         TransformMethod: answers.TransformMethod,
//         LoadDestination: answers.LoadDestination,
//         Dependencies:    determineDependencies(answers.ExtractMethod, answers.TransformMethod, answers.LoadDestination),
//     }
    
//     // Create project directory
//     if err := os.MkdirAll(answers.ProjectName, 0755); err != nil {
//         return fmt.Errorf("failed to create project directory: %w", err)
//     }
    
//     // Generate the project files
//     if err := generateProjectFiles(answers.ProjectName, etlData); err != nil {
//         return err
//     }
    
//     // Initialize virtual environment if requested
//     if answers.CreateVenv {
//         fmt.Println("\n🐍 Initializing virtual environment...")
//         if err := initializeVirtualEnv(answers.ProjectName); err != nil {
//             return err
//         }
//     }
    
//     fmt.Printf("\n✅ Project successfully generated in %s/\n", answers.ProjectName)
//     return nil
// }

// ... existing code ...

// promptETLProjectDetails collects details for an ETL project
func promptETLProjectDetails() error {
    // Questions for ETL project
    questions := []*survey.Question{
        {
            Name: "projectName",
            Prompt: &survey.Input{
                Message: "Project name:",
                Default: "python-etl-project",
                Help:    "The name of your project directory and package",
            },
            Validate: survey.Required,
        },
        {
            Name: "extractMethod",
            Prompt: &survey.Select{
                Message: "Select extraction method:",
                Options: []string{"file", "api", "database"},
                Default: "file",
                Description: func(value string, index int) string {
                    switch value {
                    case "file":
                        return "Extract data from CSV, Excel, or other files"
                    case "api":
                        return "Extract data from REST APIs"
                    case "database":
                        return "Extract data from SQL databases"
                    default:
                        return ""
                    }
                },
            },
        },
        {
            Name: "transformMethod",
            Prompt: &survey.Select{
                Message: "Select transformation method:",
                Options: []string{"basic", "advanced"},
                Default: "basic",
                Description: func(value string, index int) string {
                    switch value {
                    case "basic":
                        return "Simple data cleaning and formatting"
                    case "advanced":
                        return "Advanced processing including feature engineering, scaling, etc."
                    default:
                        return ""
                    }
                },
            },
        },
        {
            Name: "loadDestination",
            Prompt: &survey.Select{
                Message: "Select load destination:",
                Options: []string{"file", "database", "api"},
                Default: "file",
                Description: func(value string, index int) string {
                    switch value {
                    case "file":
                        return "Load data to CSV, Excel, or other files"
                    case "database":
                        return "Load data to SQL databases"
                    case "api":
                        return "Load data to REST APIs"
                    default:
                        return ""
                    }
                },
            },
        },
        {
            Name: "createVenv",
            Prompt: &survey.Confirm{
                Message: "Initialize a virtual environment?",
                Default: false,
                Help:    "Creates a Python virtual environment and installs dependencies",
            },
        },
    }
    
    // Store answers
    answers := struct {
        ProjectName     string
        ExtractMethod   string
        TransformMethod string
        LoadDestination string
        CreateVenv      bool
    }{}
    
    // Ask the questions
    err := survey.Ask(questions, &answers)
    if err != nil {
        return err
    }
    
    // ------ADD STEP 4 HERE: Advanced Dialogs based on choices------
    
    // Configuration details based on extract method
    var extractConfig struct {
        SourceType string
        Connection map[string]string
    }
    extractConfig.Connection = make(map[string]string)
    
    switch answers.ExtractMethod {
    case "file":
        var fileType string
        filePrompt := &survey.Select{
            Message: "What type of files will you extract from?",
            Options: []string{"CSV", "Excel", "JSON", "Parquet", "Other"},
            Default: "CSV",
        }
        survey.AskOne(filePrompt, &fileType)
        extractConfig.SourceType = fileType
        
        // Ask for file path pattern
        var filePattern string
        patternPrompt := &survey.Input{
            Message: "File path or pattern to extract from:",
            Default: "data/*.csv",
            Help:    "Path or glob pattern for input files (e.g., data/*.csv)",
        }
        survey.AskOne(patternPrompt, &filePattern)
        extractConfig.Connection["pattern"] = filePattern
        
    case "api":
        var apiType string
        apiPrompt := &survey.Select{
            Message: "What type of API?",
            Options: []string{"REST", "GraphQL", "SOAP", "Other"},
            Default: "REST",
        }
        survey.AskOne(apiPrompt, &apiType)
        extractConfig.SourceType = apiType
        
        // Ask for API URL
        var apiURL string
        urlPrompt := &survey.Input{
            Message: "Base API URL:",
            Default: "https://api.example.com/data",
            Help:    "The base URL for the API you'll extract from",
        }
        survey.AskOne(urlPrompt, &apiURL)
        extractConfig.Connection["url"] = apiURL
        
        // Ask if authentication is needed
        var needsAuth bool
        authPrompt := &survey.Confirm{
            Message: "Does this API require authentication?",
            Default: true,
        }
        survey.AskOne(authPrompt, &needsAuth)
        
        if needsAuth {
            var authType string
            authTypePrompt := &survey.Select{
                Message: "Authentication type:",
                Options: []string{"API Key", "OAuth2", "Basic Auth", "Bearer Token", "None"},
                Default: "API Key",
            }
            survey.AskOne(authTypePrompt, &authType)
            extractConfig.Connection["auth_type"] = authType
        }
        
    case "database":
        var dbType string
        dbPrompt := &survey.Select{
            Message: "What database will you extract from?",
            Options: []string{"PostgreSQL", "MySQL", "SQLite", "Oracle", "SQL Server", "Other"},
            Default: "PostgreSQL",
        }
        survey.AskOne(dbPrompt, &dbType)
        extractConfig.SourceType = dbType
        
        // Ask for connection details if not SQLite
        if dbType != "SQLite" {
            var host string
            hostPrompt := &survey.Input{
                Message: "Database host:",
                Default: "localhost",
            }
            survey.AskOne(hostPrompt, &host)
            extractConfig.Connection["host"] = host
            
            var port string
            portDefault := "5432" // Default for PostgreSQL
            if dbType == "MySQL" {
                portDefault = "3306"
            } else if dbType == "SQL Server" {
                portDefault = "1433"
            } else if dbType == "Oracle" {
                portDefault = "1521"
            }
            
            portPrompt := &survey.Input{
                Message: "Database port:",
                Default: portDefault,
            }
            survey.AskOne(portPrompt, &port)
            extractConfig.Connection["port"] = port
            
            var dbName string
            dbNamePrompt := &survey.Input{
                Message: "Database name:",
                Default: "mydatabase",
            }
            survey.AskOne(dbNamePrompt, &dbName)
            extractConfig.Connection["database"] = dbName
        } else {
            // For SQLite, just ask for the file path
            var dbPath string
            pathPrompt := &survey.Input{
                Message: "SQLite database file path:",
                Default: "database.sqlite",
            }
            survey.AskOne(pathPrompt, &dbPath)
            extractConfig.Connection["path"] = dbPath
        }
    }
    
    // Configuration details based on load destination
    var loadConfig struct {
        DestType  string
        Connection map[string]string
    }
    loadConfig.Connection = make(map[string]string)
    
    switch answers.LoadDestination {
    case "file":
        var fileType string
        filePrompt := &survey.Select{
            Message: "What type of files will you save to?",
            Options: []string{"CSV", "Excel", "JSON", "Parquet", "Other"},
            Default: "CSV",
        }
        survey.AskOne(filePrompt, &fileType)
        loadConfig.DestType = fileType
        
        // Ask for output directory
        var outputDir string
        dirPrompt := &survey.Input{
            Message: "Output directory:",
            Default: "output/",
            Help:    "Directory where output files will be saved",
        }
        survey.AskOne(dirPrompt, &outputDir)
        loadConfig.Connection["output_dir"] = outputDir
    
    case "database":
        // Reuse the database type selection logic from extract
        if answers.ExtractMethod == "database" && extractConfig.SourceType != "" {
            // Ask if using same database as extract
            var sameDB bool
            sameDBPrompt := &survey.Confirm{
                Message: "Use same database connection as extract?",
                Default: true,
            }
            survey.AskOne(sameDBPrompt, &sameDB)
            
            if sameDB {
                loadConfig.DestType = extractConfig.SourceType
                loadConfig.Connection = extractConfig.Connection
                
                // Just ask for the table name
                var tableName string
                tablePrompt := &survey.Input{
                    Message: "Output table name:",
                    Default: "etl_output",
                }
                survey.AskOne(tablePrompt, &tableName)
                loadConfig.Connection["table"] = tableName
                
                break
            }
        }
        
        // If not using same DB, ask for DB type and connection details
        var dbType string
        dbPrompt := &survey.Select{
            Message: "What database will you load to?",
            Options: []string{"PostgreSQL", "MySQL", "SQLite", "Oracle", "SQL Server", "Other"},
            Default: "PostgreSQL",
        }
        survey.AskOne(dbPrompt, &dbType)
        loadConfig.DestType = dbType
        
        // Ask for connection details (similar to extract)
        // ... (same code as in the extract section for database connections)
    
    case "api":
        var apiURL string
        urlPrompt := &survey.Input{
            Message: "API endpoint URL:",
            Default: "https://api.example.com/upload",
            Help:    "The URL where data will be sent",
        }
        survey.AskOne(urlPrompt, &apiURL)
        loadConfig.Connection["url"] = apiURL
        
        var method string
        methodPrompt := &survey.Select{
            Message: "HTTP method:",
            Options: []string{"POST", "PUT", "PATCH"},
            Default: "POST",
        }
        survey.AskOne(methodPrompt, &method)
        loadConfig.Connection["method"] = method
    }
    
    // ------ADD STEP 5 HERE: Multiselect for Dependencies------
    
    // Get base dependencies
    dependencies := determineDependencies(answers.ExtractMethod, answers.TransformMethod, answers.LoadDestination)
    
    // Ask for additional dependencies
    additionalDeps := []string{}
    depsPrompt := &survey.MultiSelect{
        Message: "Select additional dependencies to include:",
        Options: []string{
            "matplotlib - Plotting library for data visualization",
            "seaborn - Statistical data visualization",
            "pytest-cov - Test coverage for pytest",
            "black - Python code formatter",
            "mypy - Static type checker",
            "pydantic - Data validation and settings management",
            "fastapi - Modern, fast web framework",
            "gunicorn - WSGI HTTP Server for UNIX",
            "dash - Interactive web-based dashboards",
            "prefect - Workflow management system",
            "dask - Parallel computing library",
            "joblib - Pipeline parallelization",
            "xlrd - Excel file reading",
            "openpyxl - Excel file reading/writing",
            "statsmodels - Statistical modeling",
            "scipy - Scientific computing",
            "pyarrow - Apache Arrow data processing",
            "boto3 - AWS SDK for Python",
            "azure-storage-blob - Azure Blob Storage",
            "google-cloud-storage - Google Cloud Storage",
        },
        Help: "Use space to select, enter to confirm",
    }
    survey.AskOne(depsPrompt, &additionalDeps)
    
    // Process the selected dependencies
    for _, dep := range additionalDeps {
        // Extract just the package name before the dash
        packageName := strings.Split(dep, " - ")[0]
        dependencies = append(dependencies, packageName)
    }
    
    // Show summary
    fmt.Println("\n📋 Project Summary:")
    fmt.Printf("  • Name: %s\n", answers.ProjectName)
    fmt.Printf("  • Extract: %s (%s)\n", answers.ExtractMethod, extractConfig.SourceType)
    fmt.Printf("  • Transform: %s\n", answers.TransformMethod)
    fmt.Printf("  • Load: %s (%s)\n", answers.LoadDestination, loadConfig.DestType)
    fmt.Printf("  • Virtual Environment: %v\n", answers.CreateVenv)
    fmt.Printf("  • Dependencies: %d packages\n", len(dependencies))
    
    // Confirm generation
    proceed := false
    prompt := &survey.Confirm{
        Message: "Generate this project?",
        Default: true,
    }
    survey.AskOne(prompt, &proceed)
    
    if !proceed {
        fmt.Println("Project generation cancelled.")
        return nil
    }
    
    // Generate the template
    fmt.Println("\n🔨 Generating project...")
    
    // Use existing template generator with the enhanced data
    etlData := ETLTemplateData{
        TemplateData: TemplateData{
            ProjectName:   answers.ProjectName,
            PackageName:   strings.ReplaceAll(strings.ToLower(answers.ProjectName), "-", "_"),
            Description:   "A Python ETL project for data processing",
            PythonVersion: ">=3.8",
        },
        ExtractMethod:   answers.ExtractMethod,
        TransformMethod: answers.TransformMethod,
        LoadDestination: answers.LoadDestination,
        Dependencies:    dependencies,
    }
    
    // You may want to add the configurations to the template data
    // either by extending the ETLTemplateData struct or by adding them to a map
    
    // Create project directory
    if err := os.MkdirAll(answers.ProjectName, 0755); err != nil {
        return fmt.Errorf("failed to create project directory: %w", err)
    }
    
    // Generate the project files
    if err := generateProjectFiles(answers.ProjectName, etlData); err != nil {
        return err
    }
    
    // Initialize virtual environment if requested
    if answers.CreateVenv {
        fmt.Println("\n🐍 Initializing virtual environment...")
        if err := initializeVirtualEnv(answers.ProjectName); err != nil {
            return err
        }
    }
    
    fmt.Printf("\n✅ Project successfully generated in %s/\n", answers.ProjectName)
    fmt.Println("\n🚀 Next steps:")
    fmt.Printf("  cd %s\n", answers.ProjectName)
    if answers.CreateVenv {
        if runtime.GOOS == "windows" {
            fmt.Println("  venv\\Scripts\\activate")
        } else {
            fmt.Println("  source venv/bin/activate")
        }
    } else {
        fmt.Println("  python -m venv venv")
        if runtime.GOOS == "windows" {
            fmt.Println("  venv\\Scripts\\activate")
        } else {
            fmt.Println("  source venv/bin/activate")
        }
        fmt.Println("  pip install -r requirements.txt")
    }
    fmt.Println("  python -m src.main")
    
    return nil
}

