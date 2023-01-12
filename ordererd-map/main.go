package main

import (
	"fmt"
)

type RoleMap struct {
	M    map[string]string
	Keys []string
}

func NewRoleMap() *RoleMap {
	return &RoleMap{
		M: make(map[string]string),
	}
}

func (n *RoleMap) Set(k string, v string) {
	n.M[k] = v
	n.Keys = append(n.Keys, k)
}

func main() {
	withoutMapSlice()
	withMapSlice()
}

func withoutMapSlice() {
	roleMap := make(map[string]string)

	roleMap["support"] = "Crystal Maiden"
	roleMap["support-2"] = "Tigreal"
	roleMap["offlaner"] = "Tuskar"
	roleMap["midlaner"] = "Sniper"
	roleMap["carry"] = "Juggernaut"

	fmt.Println("Without Map Slice")
	for role, name := range roleMap {
		fmt.Printf("Role: %s Hero: %s \n", role, name)
	}
}

func withMapSlice() {

	roleMap := NewRoleMap()

	roleMap.Set("support", "Crystal Maiden")
	roleMap.Set("support-2", "Tigreal")
	roleMap.Set("offlaner", "Tuskar")
	roleMap.Set("midlaner", "Sniper")
	roleMap.Set("carry", "Juggernaut")

	fmt.Println("With Map Slice")

	// printing with the key name
	for _, role := range roleMap.Keys {
		name := roleMap.M[role]
		fmt.Printf("Role: %s Hero: %s \n", role, name)
	}
}
