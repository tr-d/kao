package main

// TODO packed kaomoji data

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	all := flag.Bool("all", false, "show all")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, getUsage())
		os.Exit(64)
	}
	kk, ok := kao[flag.Arg(0)]
	if !ok {
		os.Exit(2)
	}
	keys := []string{}
	for k := range kk {
		keys = append(keys, k)
		if *all {
			fmt.Println(k)
		}
	}
	if *all {
		os.Exit(0)
	}
	fmt.Println(keys[rand.Intn(len(keys))])
}

func getUsage() string {
	u := `NAME
	get random kaomoji by keyword

SYNOPSIS
        %s [-all] 顔類

DESCRIPTION
        return a random 顔文字 scoped by a keyword 顔類
        the -all flag prints all 顔文字 for a given 顔類

顔類 KAORUI "FACE TYPE"
        valid 顔類:
          - happy
          - sad
          - angry
          - evs
          - suss
          - hello
          - magic

EXIT CODE
        2  incorrect 顔類
        64 usage error

SOURCES
        kaomoji.ru
`
	return fmt.Sprintf(u, os.Args[0])
}

var kao = map[string]map[string]bool{
	"happy": map[string]bool{"(´｡• ᵕ •｡`)": true, "(* ^ ω ^)": true, "(´ ∀ ` *)": true, "٩(◕‿◕｡)۶": true, "☆*:.｡.o(≧▽≦)o.｡.:*☆": true, "(o^▽^o)": true, "(⌒▽⌒)☆": true, "<(￣︶￣)>": true, "。.:☆*:･'(*⌒―⌒*)))": true, "ヽ(・∀・)ﾉ": true, "(´｡• ω •｡`)": true, "(￣ω￣)": true, "｀;:゛;｀;･(°ε° )": true, "(o･ω･o)": true, "(＠＾◡＾)": true, "ヽ(*・ω・)ﾉ": true, "(o_ _)ﾉ彡☆": true, "(^人^)": true, "(o´▽`o)": true, "(*´▽`*)": true, "｡ﾟ( ﾟ^∀^ﾟ)ﾟ｡": true, "( ´ ω ` )": true, "(((o(*°▽°*)o)))": true, "(≧◡≦)": true, "(o´∀`o)": true, "(´• ω •`)": true, "(＾▽＾)": true, "(⌒ω⌒)": true, "∑d(°∀°d)": true, "╰(▔∀▔)╯": true, "(─‿‿─)": true, "(*^‿^*)": true, "ヽ(o^ ^o)ﾉ": true, "(✯◡✯)": true, "(◕‿◕)": true, "(*≧ω≦*)": true, "(☆▽☆)": true, "(⌒‿⌒)": true, "＼(≧▽≦)／": true, "ヽ(o＾▽＾o)ノ": true, "☆ ～('▽^人)": true, "(*°▽°*)": true, "٩(｡•́‿•̀｡)۶": true, "(✧ω✧)": true, "ヽ(*⌒▽⌒*)ﾉ": true, "( ´ ▽ ` )": true, "(￣▽￣)": true, "╰(*´︶`*)╯": true, "ヽ(>∀<☆)ノ": true, "o(≧▽≦)o": true, "(☆ω☆)": true, "(っ˘ω˘ς )": true, "＼(￣▽￣)／": true, "(*¯︶¯*)": true, "＼(＾▽＾)／": true, "٩(◕‿◕)۶": true, "(o˘◡˘o)": true, `\(★ω★)/`: true, `\(^ヮ^)/`: true, "(〃＾▽＾〃)": true, "(╯✧▽✧)╯": true, "o(>ω<)o": true, "o( ❛ᴗ❛ )o": true, "｡ﾟ(TヮT)ﾟ｡": true, "( ‾́ ◡ ‾́ )": true, "(ﾉ´ヮ`)ﾉ*: ･ﾟ": true, "(b ᵔ▽ᵔ)b": true, "(๑˃ᴗ˂)ﻭ": true, "(๑˘︶˘๑)": true, "( ˙꒳​˙ )": true, "(*꒦ິ꒳꒦ີ)": true, "°˖✧◝(⁰▿⁰)◜✧˖°": true, "(´･ᴗ･ ` )": true, "(ﾉ◕ヮ◕)ﾉ*:･ﾟ✧": true, "(„• ֊ •„)": true, "(.❛ ᴗ ❛.)": true, "(⁀ᗢ⁀)": true, "(￢‿￢ )": true, "(¬‿¬ )": true},
	"sad":   map[string]bool{"(ノ_<。)": true, "(-_-)": true, "(´-ω-`)": true, ".･ﾟﾟ･(／ω＼)･ﾟﾟ･.": true, "(μ_μ)": true, "(ﾉД`)": true, "(-ω-、)": true, "。゜゜(´Ｏ`) ゜゜。": true, "o(TヘTo)": true, "( ; ω ; )": true, "(｡╯︵╰｡)": true, "｡･ﾟﾟ*(>д<)*ﾟﾟ･｡": true, "( ﾟ，_ゝ｀)": true, "(个_个)": true, "(╯︵╰,)": true, "｡･ﾟ(ﾟ><ﾟ)ﾟ･｡": true, "( ╥ω╥ )": true, "(╯_╰)": true, "(╥_╥)": true, ".｡･ﾟﾟ･(＞_＜)･ﾟﾟ･｡.": true, "(／ˍ・、)": true, "(ノ_<、)": true, "(╥﹏╥)": true, "｡ﾟ(｡ﾉωヽ｡)ﾟ｡": true, "(つω`｡)": true, "(｡T ω T｡)": true, "(ﾉω･､)": true, "･ﾟ･(｡>ω<｡)･ﾟ･": true, "(T_T)": true, "(>_<)": true, "(っ˘̩╭╮˘̩)っ": true, "｡ﾟ･ (>﹏<) ･ﾟ｡": true, "o(〒﹏〒)o": true, "(｡•́︿•̀｡)": true, "(ಥ﹏ಥ)": true},
	"angry": map[string]bool{"(＃`Д´)": true, "(`皿´＃)": true, "( ` ω ´ )": true, "ヽ( `д´*)ノ": true, "(・`ω´・)": true, "(`ー´)": true, "ヽ(`⌒´メ)ノ": true, "凸(`△´＃)": true, "( `ε´ )": true, "ψ( ` ∇ ´ )ψ": true, "ヾ(`ヘ´)ﾉﾞ": true, "ヽ(‵﹏´)ノ": true, "(ﾒ` ﾛ ´)": true, "(╬`益´)": true, "┌∩┐(◣_◢)┌∩┐": true, "凸( ` ﾛ ´ )凸": true, "Σ(▼□▼メ)": true, "(°ㅂ°╬)": true, "ψ(▼へ▼メ)～→": true, "(ノ°益°)ノ": true, "(҂ `з´ )": true, "(‡▼益▼)": true, "(҂` ﾛ ´)凸": true, "((╬◣﹏◢))": true, "٩(╬ʘ益ʘ╬)۶": true, "(╬ Ò﹏Ó)": true, "＼＼٩(๑`^´๑)۶／／": true, "(凸ಠ益ಠ)凸": true, "↑_(ΦwΦ)Ψ": true, "←~(Ψ▼ｰ▼)∈": true, "୧((#Φ益Φ#))୨": true, "٩(ఠ益ఠ)۶": true, "(ﾉಥ益ಥ)ﾉ": true},
	"evs":   map[string]bool{"ヽ(ー_ー )ノ": true, "ヽ(´ー` )┌": true, "┐(‘～` )┌": true, "ヽ(　￣д￣)ノ": true, "┐(￣ヘ￣)┌": true, "ヽ(￣～￣　)ノ": true, "╮(￣_￣)╭": true, "ヽ(ˇヘˇ)ノ": true, "┐(￣～￣)┌": true, "┐(︶▽︶)┌": true, "╮(￣～￣)╭": true, `¯\_(ツ)_/¯`: true, "┐( ´ д ` )┌": true, "╮(︶︿︶)╭": true, "┐(￣∀￣)┌": true, "┐( ˘ ､ ˘ )┌": true, "╮(︶▽︶)╭": true, "╮( ˘ ､ ˘ )╭": true, "┐( ˘_˘ )┌": true, "╮( ˘_˘ )╭": true, "┐(￣ヮ￣)┌": true, "ᕕ( ᐛ )ᕗ": true},
	"suss":  map[string]bool{"(￢_￢)": true, "(→_→)": true, "(￢ ￢)": true, "(￢‿￢ )": true, "(¬_¬ )": true, "(←_←)": true, "(¬ ¬ )": true, "(¬‿¬ )": true, "(↼_↼)": true, "(⇀_⇀)": true},
	"hello": map[string]bool{"(*・ω・)ﾉ": true, "(￣▽￣)ノ": true, "(°▽°)/": true, "( ´ ∀ ` )ﾉ": true, "(^-^*)/": true, "(＠´ー`)ﾉﾞ": true, "(´• ω •`)ﾉ": true, "( ° ∀ ° )ﾉﾞ": true, "ヾ(*'▽'*)": true, "＼(⌒▽⌒)": true, "ヾ(☆▽☆)": true, "( ´ ▽ ` )ﾉ": true, "(^０^)ノ": true, "~ヾ(・ω・)": true, "(・∀・)ノ": true, "ヾ(^ω^*)": true, "(*°ｰ°)ﾉ": true, "(・_・)ノ": true, "(o´ω`o)ﾉ": true, "ヾ(☆'∀'☆)": true, "(￣ω￣)/": true, "( ´ ω ` )ノﾞ": true, "(⌒ω⌒)ﾉ": true, "(o^ ^o)/": true, "(≧▽≦)/": true, "(✧∀✧)/": true, "(o´▽`o)ﾉ": true, "(￣▽￣)/": true},
	"magic": map[string]bool{"(ノ ˘_˘)ノ　ζ|||ζ　ζ|||ζ　ζ|||ζ": true, "(ﾉ≧∀≦)ﾉ ‥…━━━★": true, "(ﾉ>ω<)ﾉ :｡･:*:･ﾟ’★,｡･:*:･ﾟ’☆": true, "(ノ°∀°)ノ⌒･*:.｡. .｡.:*･゜ﾟ･*☆": true, "╰( ͡° ͜ʖ ͡° )つ──☆*:・ﾟ": true, "(＃￣□￣)o━∈・・━━━━☆": true, "(⊃｡•́‿•̀｡)⊃━✿✿✿✿✿✿": true, "(∩ᄑ_ᄑ)⊃━☆ﾟ*･｡*･:≡( ε:)": true, "(/￣ー￣)/~~☆’.･.･:★’.･.･:☆": true, "(∩` ﾛ ´)⊃━炎炎炎炎炎": true},
}
