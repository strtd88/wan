package cmd

import (
	"fmt"

	"wan/db/repo"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var AddProjectCmd = &cobra.Command{
	Use:   "project add",
	Short: "Add a new project with directory and metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		desc, _ := cmd.Flags().GetString("description")
		dir, _ := cmd.Flags().GetString("dir")
		status, _ := cmd.Flags().GetString("status")
		category, _ := cmd.Flags().GetString("category")
		bizShort, _ := cmd.Flags().GetString("business-value-short")
		bizLong, _ := cmd.Flags().GetString("business-value-long")
		notePath, _ := cmd.Flags().GetString("obsidian-note-path")

		db, err := ConnectDB()
		if err != nil {
			return err
		}
		defer db.Close()

		fmt.Printf("format string", bizLong, bizShort)

		repo := repo.NewProjectRepo(db)
		//err = repo.CreateProject(name, desc, dir, status, category, bizShort, bizLong)
		err = repo.CreateProject(name, desc, dir, status, category)
		if err != nil {
			return err
		}

		if notePath != "" {
			// Optional: link Obsidian note
		}

		fmt.Printf("✅ Project added: %s\n", name)
		return nil
	},
}

func init() {
	AddProjectCmd.Flags().StringP("name", "n", "", "Project name (required)")
	AddProjectCmd.MarkFlagRequired("name")
	AddProjectCmd.Flags().StringP("description", "d", "", "Project description")
	AddProjectCmd.Flags().String("dir", "", "Directory path")
	AddProjectCmd.Flags().String("status", "new", "Status (new/active/paused/complete)")
	AddProjectCmd.Flags().String("category", "devtools", "Category tag (e.g., learning, work)")
	AddProjectCmd.Flags().String("business-value-short", "", "Brief summary of business value")
	AddProjectCmd.Flags().String("business-value-long", "", "Detailed business value")
	AddProjectCmd.Flags().String("obsidian-note-path", "", "Path to linked Obsidian note")

	rootCmd.AddCommand(AddProjectCmd)
}
