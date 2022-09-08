package channelserver

import (
	"encoding/hex"
	"erupe-ce/network/mhfpacket"
)

func handleMsgMhfGetTowerInfo(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetTowerInfo)
	var data []byte
	var err error
	/*
		type:
		1 == TOWER_RANK_POINT,
		2 == GET_OWN_TOWER_SKILL
		3 == ?
		4 == TOWER_TOUHA_HISTORY
		5 = ?

		[] = type
		req
		resp

		01 1d 01 fc 00 09 [00 00 00 01] 00 00 00 02 00 00 00 00
		00 12 01 fc 00 09 01 00 00 18 0a 21 8e ad 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00 00

		01 1d 01 fc 00 0a [00 00 00 02] 00 00 00 00 00 00 00 00
		00 12 01 fc 00 0a 01 00 00 94 0a 21 8e ad 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00

		01 1d 01 ff 00 0f [00 00 00 04] 00 00 00 00 00 00 00 00
		00 12 01 ff 00 0f 01 00 00 24 0a 21 8e ad 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00

		01 1d 01 fc 00 0b [00 00 00 05] 00 00 00 00 00 00 00 00
		00 12 01 fc 00 0b 01 00 00 10 0a 21 8e ad 00 00 00 00 00 00 00 00 00 00 00 00
	*/
	switch pkt.InfoType {
	case mhfpacket.TowerInfoTypeTowerRankPoint:
		data, err = hex.DecodeString("0A218EAD0000000000000000000000010000000000000000")
	case mhfpacket.TowerInfoTypeGetOwnTowerSkill:
		//data, err = hex.DecodeString("0A218EAD000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
		data, err = hex.DecodeString("0A218EAD0000000000000000000000010000001C0000000500050000000000020000000000000000000000000000000000030003000000000003000500050000000300030003000300030003000200030001000300020002000300010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	case mhfpacket.TowerInfoTypeUnk3:
		panic("No known response values for TowerInfoTypeUnk3")
	case mhfpacket.TowerInfoTypeTowerTouhaHistory:
		data, err = hex.DecodeString("0A218EAD0000000000000000000000010000000000000000000000000000000000000000")
	case mhfpacket.TowerInfoTypeUnk5:
		data, err = hex.DecodeString("0A218EAD000000000000000000000000")
	}

	if err != nil {
		stubGetNoResults(s, pkt.AckHandle)
	}
	doAckBufSucceed(s, pkt.AckHandle, data)
}

func handleMsgMhfPostTowerInfo(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfPostTowerInfo)
	doAckSimpleSucceed(s, pkt.AckHandle, []byte{0x00, 0x00, 0x00, 0x00})
}

func handleMsgMhfGetGemInfo(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetGemInfo)
	doAckBufSucceed(s, pkt.AckHandle, make([]byte, 8))
}

func handleMsgMhfPostGemInfo(s *Session, p mhfpacket.MHFPacket) {}
