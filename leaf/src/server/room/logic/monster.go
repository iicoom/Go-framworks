package logic

import (
	sample "myLeaf/msg/MyGame/Sample"
	flatbuffers "github.com/google/flatbuffers/go"
)

// Example how to use Flatbuffers to create and read binary buffers.
func WriteToFlat() []byte{
	builder := flatbuffers.NewBuilder(0)
	// 1. 构建 table Weapon
	// Create some weapons for our Monster ("Sword" and "Axe").
	weaponOne := builder.CreateString("Sword")
	weaponTwo := builder.CreateString("Axe")

	sample.WeaponStart(builder)
	sample.WeaponAddName(builder, weaponOne)
	sample.WeaponAddDamage(builder, 3)
	sword := sample.WeaponEnd(builder)

	sample.WeaponStart(builder)
	sample.WeaponAddName(builder, weaponTwo)
	sample.WeaponAddDamage(builder, 5)
	axe := sample.WeaponEnd(builder)

	// 2. 构建 table Monster
	// Serialize the FlatBuffer data.
	name := builder.CreateString("Orc")  // Orc妖魔

	// table Monster - inventory
	sample.MonsterStartInventoryVector(builder, 10) // 库存数组的容量为10
	// Note: Since we prepend the bytes, this loop iterates in reverse.
	for i := 9; i >= 0; i-- {
		builder.PrependByte(byte(i))
	}
	inv := builder.EndVector(10)

	// table Monster - weapons
	sample.MonsterStartWeaponsVector(builder, 2)
	// Note: Since we prepend the weapons, prepend in reverse order.
	builder.PrependUOffsetT(axe)
	builder.PrependUOffsetT(sword)
	weapons := builder.EndVector(2)

	// 3. 构建 table struct Vec3
	pos := sample.CreateVec3(builder, 1.0, 2.0, 3.0)

	// 注意下面的组装顺序 与 monster 中一致 值可以覆盖
	sample.MonsterStart(builder)
	sample.MonsterAddPos(builder, pos)
	sample.MonsterAddHp(builder, 300)
	sample.MonsterAddName(builder, name)
	sample.MonsterAddInventory(builder, inv)
	sample.MonsterAddColor(builder, sample.ColorRed)
	sample.MonsterAddWeapons(builder, weapons)
	sample.MonsterAddEquippedType(builder, sample.EquipmentWeapon)
	sample.MonsterAddEquipped(builder, axe)
	orc := sample.MonsterEnd(builder)

	builder.Finish(orc)

	// We now have a FlatBuffer that we could store on disk or send over a network.
	// ...Saving to file or sending over a network code goes here...
	buf := builder.FinishedBytes()
	return buf
}

func ReadFromFlat() {
	// Instead, we are going to access this buffer right away (as if we just received it).
	//相反，我们将立即访问这个缓冲区(就像我们刚刚收到它一样)

}
