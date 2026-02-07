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
				{Title: "memo_root", Value: &now.MemoRoot},
				{Title: "article_root", Value: &now.ArticleRoot},
				{Title: "notion_token", Value: &now.NotionToken},
				{Title: "default_memo_db", Value: &now.DefaultMemoDB},
			}
			if err := prompt.UpdateInputs(inputs); err != nil {
				return err
			}

			MapInput := prompt.MapInput{
				KeyDescritpiton:   "Notion Database key > ",
				ValueDescritpiton: "Notion Database Id > ",
				QuitMark:          "q",
				Map:               now.Databases,
			}
			// Mapの初期化を入れる入れるならmapを定義する構造体もしくはマップを定義する際に利用するより共通の構造体
			// アーキとしてコモン的なものを作るのはありかどうか
			prompt.MapInputs(MapInput)

			// prettyPrint(now.Databases) // This line is now redundant
			cfg, err := Mapper.InitRequestDtoToConfig(now)
			fmt.Println(cfg, err)
			if err != nil {
				fmt.Println(err)
				return err
			}

			return configUseCase.SaveOverWrite(cfg)
		},
	}
	return initCmd
}
