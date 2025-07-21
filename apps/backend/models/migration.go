package models

func AllModels() []interface{} {
    return []interface{}{
        &User{},
        // &OtherModel{}, // 必要に応じて追加
    }
}