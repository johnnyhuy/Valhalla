package inventory

func (item *Item) Save(charID int32) error {
	// delete all items // This is inneficient as it happens for each item in inventory

	// save item

	// query := `UPDATE characters set skin=?, hair=?, face=?, level=?,
	// job=?, str=?, dex=?, intt=?, luk=?, hp=?, maxHP=?, mp=?, maxMP=?,
	// ap=?, sp=?, exp=?, fame=?, mapID=?, mesos=? WHERE id=?`

	// // need to calculate nearest spawn point for mapPos

	// records, err := connection.Db.Query(query,
	// 	char.GetSkin(),
	// 	char.GetHair(),
	// 	char.GetFace(),
	// 	char.GetLevel(),
	// 	char.GetJob(),
	// 	char.GetStr(),
	// 	char.GetDex(),
	// 	char.GetInt(),
	// 	char.GetLuk(),
	// 	char.GetHP(),
	// 	char.GetMaxHP(),
	// 	char.GetMP(),
	// 	char.GetMaxMP(),
	// 	char.GetAP(),
	// 	char.GetSP(),
	// 	char.GetEXP(),
	// 	char.GetFame(),
	// 	char.GetCurrentMap(),
	// 	char.GetMesos(),
	// 	char.GetCharID())

	// defer records.Close()

	return nil
}
