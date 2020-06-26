package dndlookup

import (
  "strings"

  "github.com/bwmarrin/discordgo"
  "golords/handlers/create/handler"
  go5e "github.com/elliotcubit/go-5e-srd-api"
)

func New() handler.CreateHandler {
  return DndLookup{}
}

type DndLookup struct {
  handler.DefaultHandler
}

func (d DndLookup) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m.Content, " ", 3)
  if len(data) < 3 {
    return
  }

  var out string
  data[2] = strings.ReplaceAll(data[2], " ", "+")
  data[1] = strings.ToLower(data[1])

  switch data[1]{
  case "abilityscore":
    out = doAbilityScore(data[2])
  case "class":
    out = doClass(data[2])
  case "condition":
    out = doCondition(data[2])
  case "damagetype":
    out = doDamageType(data[2])
  case "equipmentcategory":
    out = doEquipmentCategory(data[2])
  case "equipment":
    out = doEquipment(data[2])
  case "feature":
    out = doFeature(data[2])
  case "language":
    out = doLanguage(data[2])
  case "magicschool":
    out = doMagicSchool(data[2])
  case "monster":
    out = doMonster(data[2])
  case "proficiency":
    out = doProficiency(data[2])
  case "race":
    out = doRace(data[2])
  case "skill":
    out = doSkill(data[2])
  // case "spellcasting":
  case "spell":
    out = doSpell(data[2])
  // case "startingequipment":
  case "subclass":
    out = doSubclass(data[2])
  case "subrace":
    out = doSubrace(data[2])
  case "trait":
    out = doTrait(data[2])
  case "weaponproperty":
    out = doWeaponProperty(data[2])
  case "endpoints":
    out = strings.Join(go5e.Endpoints(), ", ")
  }

  // Any error that occurs in a handler will not be
  // able to be handled so we simply return the zero
  // value of a string from them.
  if out != "" {
    s.ChannelMessageSend(m.ChannelID, out)
  }
}

func (d DndLookup) GetPrompts() []string {
  return []string{"!lookup"}
}

func (d DndLookup) Help() string {
  return "Lookup information from the SRD.\n!lookup <type> <query>"
}

func (d DndLookup) Should(hint string) bool {
  prompts := d.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}
