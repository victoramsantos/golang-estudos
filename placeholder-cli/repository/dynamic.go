package repository

import (
	"math/rand"
	"time"
)

var _POEMS = make([]Poem, 0)

func initPoems() {
	_POEMS = append(_POEMS,
		Poem{
			Content: `Olhe, tenho uma alma muito prolixa e uso poucas palavras.
			Sou irritável e firo facilmente.
			Também sou muito calmo e perdoo logo.
			Não esqueço nunca.
			Mas há poucas coisas de que eu me lembre.`,
			Author: "Clarice Linspector",
		},
	)

	_POEMS = append(_POEMS,
		Poem{
			Content: `Que minha solidão me sirva de companhia.
			que eu tenha a coragem de me enfrentar.
			que eu saiba ficar com o nada
			e mesmo assim me sentir
			como se estivesse plena de tudo.`,
			Author: "Clarice Linspector",
		},
	)

	_POEMS = append(_POEMS,
		Poem{
			Content: `Motivo

			Eu canto porque o instante existe
			e a minha vida está completa.
			Não sou alegre nem sou triste:
			sou poeta.
			
			Irmão das coisas fugidias,
			não sinto gozo nem tormento.
			Atravesso noites e dias
			no vento.
			
			Se desmorono ou se edifico,
			se permaneço ou me desfaço,
			— não sei, não sei. Não sei se fico
			ou passo.
			
			Sei que canto. E a canção é tudo.
			Tem sangue eterno a asa ritmada.
			E um dia sei que estarei mudo:
			— mais nada.`,
			Author: "Cecília Meireles",
		},
	)

	_POEMS = append(_POEMS,
		Poem{
			Content: `Eu canto porque o instante existe e a minha vida está completa. 
			Não sou alegre nem sou triste: sou poeta`,
			Author: "Cecília Meireles",
		},
	)

	_POEMS = append(_POEMS,
		Poem{
			Content: `Toda a poesia - e a canção é uma poesia ajudada - reflete o que a alma não tem. 
			Por isso a canção dos povos tristes é alegre e a canção dos povos alegres é triste.`,
			Author: "Fernando Pessoa",
		},
	)
}

func sanitizePoems() {
	if len(_POEMS) == 0 {
		initPoems()
	}
}

func GetDynamicContent() Poem {
	sanitizePoems()
	randomizer := sanitizeRand()
	var randPos = randomizer.Intn(len(_POEMS))
	return _POEMS[randPos]
}

func sanitizeRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
