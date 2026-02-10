/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Melbeck777/noq/internal/application/usecase"
	"github.com/Melbeck777/noq/internal/presentation/mapper"
	"github.com/Melbeck777/noq/internal/presentation/prompt"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

// initCmd represents the init command

func newInitCmd(ui *rwi.RWI, configUseCase usecase.ConfigUseCase) *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "init memo and article",
		Long:  `setting the your memo_root, article_root and default_memo_db`,
		RunE: func(cmd *cobra.Command, args []string) error {
			currentConfig, err := configUseCase.Load()
			if err != nil {
				tmp, err := configUseCase.DefaultValue()
				if err != nil {
					return fmt.Errorf("")
				}
				currentConfig = tmp
			}
			Mapper := mapper.NewConfigMapper()
			now, err := Mapper.EntityConfigToInitRequestDto(currentConfig)
			if err != nil {
				return err
			}
			inputs := []prompt.Input{
				{Title: "memo_root", Value: &now.MemoRoot, Description: "Input full path "},
				{Title: "article_root", Value: &now.ArticleRoot, Description: "Input full path "},
				{Title: "notion_token", Value: &now.NotionToken, Description: "Input Notion's token "},
			}
			if err := prompt.UpdateInputs(inputs); err != nil {
				return err
			}

			MapInput := prompt.MapInput{
				KeyDescription:   "Notion Database key > ",
				ValueDescription: "Notion Database Id > ",
				QuitMark:         "q",
				Map:              now.Databases,
			}
			// DBにkey,valueを入力する
			prompt.MapInputs(MapInput)
			MapKey := prompt.MapKey{
				Map:   now.Databases,
				Title: "Key",
				Value: &now.DefaultMemoDB,
			}

			if err := prompt.MapKeySelect(MapKey); err != nil {
				return err
			}

			cfg, err := Mapper.InitRequestDtoToConfig(now)
			fmt.Println(cfg, err)
			if err != nil {
				fmt.Println(err)
				return err
			}
			// 上書き保存
			return configUseCase.SaveOverWrite(cfg)
		},
	}
	return initCmd
}
