package board

import (
	"testing"

	"github.com/tasmanianfox/dingo/common"
)

func TestGetAttackMap(t *testing.T) {
	p := FenToPosition("1r1r2k1/2q2ppp/b3pn2/p1bp2B1/5P1P/2NB2P1/PPPQ4/2K3RR b - - 0 1")
	am := GetAttackMap(p, common.ColourBlack)
	eam := [common.NumRows][common.NumColumns]bool{
		{false, false, false, false, false, false, true, false},
		{false, true, false, false, false, true, false, false},
		{true, true, false, true, true, false, false, false},
		{false, true, true, true, true, true, true, false},
		{true, true, true, true, true, true, false, true},
		{false, true, true, true, true, true, true, true},
		{true, true, false, true, true, true, true, true},
		{true, true, true, true, true, true, true, true},
	}
	if am != eam {
		t.Errorf("Incorrect attack map %v %s", am, PositionToFen(p))
	}
	p = FenToPosition("1r1r2k1/2q2ppp/b3pn2/p1bp2B1/5P1P/2NB2P1/PPPQ4/2K3RR w - - 0 1")
	am = GetAttackMap(p, common.ColourWhite)
	eam = [common.NumRows][common.NumColumns]bool{
		{false, true, true, true, true, true, true, true},
		{true, true, true, true, true, true, true, true},
		{true, true, true, true, true, false, true, true},
		{true, false, true, false, true, true, false, true},
		{false, true, false, true, true, true, true, false},
		{true, false, false, false, false, true, true, true},
		{false, false, false, false, false, false, false, true},
		{false, false, false, false, false, false, false, false},
	}
	if am != eam {
		t.Errorf("Incorrect attack map")
	}
}

func TestGetKingAttackMap(t *testing.T) {
	am := getKingAttackMap(common.Row1, common.ColumnA)
	eam := [common.NumRows][common.NumColumns]bool{
		{false, true, false, false, false, false, false, false},
		{true, true, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}
	if am != eam {
		t.Errorf("Incorrect attack map")
	}

	am = getKingAttackMap(common.Row1, common.ColumnE)
	eam = [common.NumRows][common.NumColumns]bool{
		{false, false, false, true, false, true, false, false},
		{false, false, false, true, true, true, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}
	if am != eam {
		t.Errorf("Incorrect attack map")
	}

	am = getKingAttackMap(common.Row4, common.ColumnE)
	eam = [common.NumRows][common.NumColumns]bool{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, true, true, true, false, false},
		{false, false, false, true, false, true, false, false},
		{false, false, false, true, true, true, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}
	if am != eam {
		t.Errorf("Incorrect attack map")
	}
}

func TestGetQueenAttackMap(t *testing.T) {
	p := FenToPosition("5k2/8/8/5r2/8/8/1PQ2P2/1BRRK3 w - - 0 1")
	am := getQueenAttackMap(p.Board, common.Row2, common.ColumnC)
	eam := [common.NumRows][common.NumColumns]bool{
		{false, true, true, true, false, false, false, false},
		{false, true, false, true, true, true, false, false},
		{false, true, true, true, false, false, false, false},
		{true, false, true, false, true, false, false, false},
		{false, false, true, false, false, true, false, false},
		{false, false, true, false, false, false, false, false},
		{false, false, true, false, false, false, false, false},
		{false, false, true, false, false, false, false, false},
	}
	if am != eam {
		t.Errorf("Incorrect attack map")
	}
}

func TestGetRookAttackMap(t *testing.T) {
	p := FenToPosition("k7/4p3/8/8/3Nr1P1/8/8/K7 b - - 0 1")
	am := getRookAttackMap(p.Board, common.Row4, common.ColumnE)
	eam := [common.NumRows][common.NumColumns]bool{
		{false, false, false, false, true, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, false, true, false, true, true, false},
		{false, false, false, false, true, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, false, false, false, false, false, false},
	}
	if am != eam {
		t.Errorf("Incorrect attack map")
	}
}

func TestGetBishopAttackMap(t *testing.T) {
	p := FenToPosition("5K1k/8/5p2/8/7B/8/8/8 w - - 0 1")
	am := getBishopAttackMap(p.Board, common.Row4, common.ColumnH)
	eam := [common.NumRows][common.NumColumns]bool{
		{false, false, false, false, true, false, false, false},
		{false, false, false, false, false, true, false, false},
		{false, false, false, false, false, false, true, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false},
		{false, false, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}
	if am != eam {
		t.Errorf("Incorrect attack map")
	}
}

func TestGetKnightAttackMap(t *testing.T) {
	am := getKnightAttackMap(common.Row7, common.ColumnF)
	eam := [common.NumRows][common.NumColumns]bool{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, true, false, true, false},
		{false, false, false, true, false, false, false, true},
		{false, false, false, false, false, false, false, false},
		{false, false, false, true, false, false, false, true},
	}
	if am != eam {
		t.Errorf("Incorrect attack map")
	}
}

func TestGetPawntAttackMap(t *testing.T) {
	am := getPawnAttackMap(common.Row5, common.ColumnE, common.ColourWhite)
	eam := [common.NumRows][common.NumColumns]bool{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, true, false, true, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}
	if am != eam {
		t.Errorf("Incorrect attack map")
	}

	am = getPawnAttackMap(common.Row2, common.ColumnA, common.ColourBlack)
	eam = [common.NumRows][common.NumColumns]bool{
		{false, true, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}
	if am != eam {
		t.Errorf("Incorrect attack map")
	}
}
