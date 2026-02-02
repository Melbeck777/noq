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
			// 入力を受け取るようにする
			// 受け取った入力をentityに押し込める
			// request dtoが必要でこれは標準の型でしゅつりょくするような形式とする
			// 入力を受付の終了としてqがおされたら入力を終了して次に進む
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
			}
			if err := prompt.UpdateInputs(inputs); err != nil {
				return err
			}
			fmt.Printf("%s, %s", now.MemoRoot, now.ArticleRoot)

			return nil
		},
	}
	return initCmd
}
