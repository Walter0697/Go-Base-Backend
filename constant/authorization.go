package constant

//for account userrole, we use bit to identify which userrole you are at
//by shifting the bit, we can get different userrole
//such as:
//0000 0001 superadmin
//0000 0010 admin
//0000 0100 normal user

//then, when checking if they are authorized, we can use bitwise or
//such as:
//0000 0011 << only admin or above is authorized
//0000 0100 | 0000 0011 => 0000 0111, not equals to 0000 0011 (original value)
//0000 0010 | 0000 0011 => 0000 0011, equals to 0000 0011 (original value)

func GetAdminAccount() uint {
	return 1
}

func GetStandardAccount() uint {
	return 1 << 1
}

func GetAdminRole() uint {
	return GetAdminAccount()
}

func GetStandardRole() uint {
	return GetStandardAccount() & GetAdminAccount()
}