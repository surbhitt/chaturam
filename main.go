package main

import (
	"fmt"
)

type Coord struct {
	x, y int
}

type Piece struct {
	color     string
	name      string
	position  Coord
	has_moved bool
	repr      rune
}

func get_char_repr(piece_name string) rune {
	switch piece_name {
	case "king":
		return 'k'
	case "queen":
		return 'q'
	case "rook":
		return 'r'
	case "bishop":
		return 'b'
	case "knight":
		return 'n'
	case "pawn":
		return 'p'
	default:
		return '#'
	}
}

func make_piece(color string, name string, position Coord) Piece {
	repr := get_char_repr(name)
	has_moved := false
	return Piece{
		color:     color,
		name:      name,
		position:  position,
		has_moved: has_moved,
		repr:      repr,
	}
}

func check_board_bound(pos Coord) bool {
	return pos.x < 8 && pos.y < 8 && pos.x > -1 && pos.y > -1
}

func get_king_moves(pos Coord) []Coord {
	var king_moves []Coord
	// create a square around the pos
	// all coords in that square except for the original pos
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if !(i == 0 && j == 0) {
				move := Coord{pos.x + i, pos.y + j}
				if check_board_bound(move) {
					king_moves = append(king_moves, move)
				}
			}
		}
	}
	return king_moves
}

func get_queen_moves(pos Coord) []Coord {
	var queen_moves []Coord
	// extrapolate diagonally, vertically and horizontally
	for i := -8; i < 8; i++ {
		move := Coord{pos.x + 0, pos.y + i}
		if check_board_bound(move) {
			queen_moves = append(queen_moves, move)
		}
		move = Coord{pos.x + i, pos.y + 0}
		if check_board_bound(move) {
			queen_moves = append(queen_moves, move)
		}
		move = Coord{pos.x + i, pos.y + i}
		if check_board_bound(move) {
			queen_moves = append(queen_moves, move)
		}
		move = Coord{pos.x + i, pos.y - i}
		if check_board_bound(move) {
			queen_moves = append(queen_moves, move)
		}
	}
	return queen_moves
}

func get_rook_moves(pos Coord) []Coord {
	var rook_moves []Coord
	// extrapolate vertically and horizontally
	for i := -8; i < 8; i++ {
		move := Coord{pos.x + 0, pos.y + i}
		if check_board_bound(move) {
			rook_moves = append(rook_moves, move)
		}
		move = Coord{pos.x + i, pos.y + 0}
		if check_board_bound(move) {
			rook_moves = append(rook_moves, move)
		}
	}
	return rook_moves
}

func get_bishop_moves(pos Coord) []Coord {
	var bishop_moves []Coord
	// extrapolate diagonally
	for i := -8; i < 8; i++ {
		move := Coord{pos.x + i, pos.y + i}
		if check_board_bound(move) {
			bishop_moves = append(bishop_moves, move)
		}
		move = Coord{pos.x - i, pos.y + i}
		if check_board_bound(move) {
			bishop_moves = append(bishop_moves, move)
		}
	}
	return bishop_moves
}

func get_knight_moves(pos Coord) []Coord {
	var knight_moves []Coord
	possible_moves := [][]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {-1, 2}, {1, 2}, {-1, -2}, {1, -2}}
	for i := 0; i < len(possible_moves); i++ {
		move := Coord{pos.x + possible_moves[i][0], pos.y + possible_moves[i][1]}
		if check_board_bound(move) {
			knight_moves = append(knight_moves, move)
		}
	}
	return knight_moves
}

func get_pawn_moves(pos Coord) []Coord {
	// TODO: check for first move and allow 2 steps
	// TODO: cross attack move
	move := Coord{pos.x - 1, pos.y}
	if check_board_bound(move) {
		return []Coord{move}
	} else {
		return []Coord{move}
	}
}

func (p Piece) get_valid_moves() []Coord {
	switch p.name {
	case "king":
		return get_king_moves(p.position)
	case "queen":
		return get_queen_moves(p.position)
	case "rook":
		return get_rook_moves(p.position)
	case "bishop":
		return get_bishop_moves(p.position)
	case "knight":
		return get_knight_moves(p.position)
	case "pawn":
		return get_pawn_moves(p.position)
	default:
		return []Coord{}
	}
}

// func initiate_board() {}
func print_piece_board(p Piece) {
	valid_moves := p.get_valid_moves()

	var board [8][8]rune

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = '-'
		}
	}

	for coord_i := 0; coord_i < len(valid_moves); coord_i++ {
		x := valid_moves[coord_i].x
		y := valid_moves[coord_i].y
		board[x][y] = '0'
	}

	board[p.position.x][p.position.y] = p.repr

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

// func add_pieces(board []Piece, color string) {
// }

func add_pawns(board *[8][8]Piece, color string) {
	if color == "black" {
		for i := 0; i < 8; i++ {
			board[1][i] = make_piece(color, "pawn", Coord{1, i})
		}
	}
	if color == "white" {
		for i := 0; i < 8; i++ {
			board[6][i] = make_piece(color, "pawn", Coord{6, i})
		}
	}
}

func initiate_board() [8][8]Piece {
	var board [8][8]Piece
	add_pawns(&board, "white")
	add_pawns(&board, "black")
	// add_pieces()
	// add_pieces()
	// for i := 0; i < 8; i++ {
	// 	for j := 0; j < 8; j++ {
	// board[i][j] = nil
	// 	}
	// }
	return board
}

func print_board(board [8][8]Piece) {
	var board_rep [8][8]rune

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
            board_rep[i][j] = '-'
        }
    }
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
            if (board[i][j] != Piece{}) {
                board_rep[i][j] = board[i][j].repr
            }
		}
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%c ", board_rep[i][j])
		}
		fmt.Println()
	}
}

func main() {
	board := initiate_board()
	// p := make_piece("white", "knight", coord{x: 4, y: 4})
	print_board(board)
}
