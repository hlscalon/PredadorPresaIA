package main

import (
	//"log"
)

const VelocidadeMaximaPresa = 2
const QtdeIteracaoPanico = 10

type Presa struct {
	AgenteImpl
	morreu bool
	qualidadeEmocao int
	intensidadeEmocao int
	mudouCor bool
	cAgente CAgente
	qtdeIteracaoPanico int
}

func (p *Presa) getCAgente() CAgente {
	if p.mudouCor == true {
		return C_Presa_Fugindo
	}
	return C_Presa
}

func (p *Presa) alterarQualidadeEmocao(valor, limiteInferior, limiteSuperior int) {
	if limiteInferior < -3 {
		limiteInferior = -3
	}

	if limiteSuperior > 3 {
		limiteSuperior = 3
	}

	p.qualidadeEmocao += valor
	if p.qualidadeEmocao < limiteInferior {
		p.qualidadeEmocao = limiteInferior
	} else if p.qualidadeEmocao > limiteSuperior {
		p.qualidadeEmocao = limiteSuperior
	}
}

func (p *Presa) alterarIntensidadeEmocao(valor int, limiteInferior, limiteSuperior int) {
	if limiteInferior < 0 {
		limiteInferior = 0
	}

	if limiteSuperior > 3 {
		limiteSuperior = 3
	}

	p.intensidadeEmocao += valor
	if p.intensidadeEmocao < limiteInferior {
		p.intensidadeEmocao = limiteInferior
	} else if p.intensidadeEmocao > limiteSuperior {
		p.intensidadeEmocao = limiteSuperior
	}
}

func (p *Presa) mover(campoVisao CampoVisao) (Posicao, PosMovimento) {
	// verifica se tem predador
	// verifica se tem presa que mudou de cor (???)

	direcaoPredador := Direcao(-1)
	direcaoPresaFugindo := Direcao(-1)
	qtdePredadores := 0
	qtdePresasLivres := 0
	qtdePresasFugindo := 0
	for i, campo := range campoVisao.Posicoes {
		if campo.Agente == C_Predador {
			direcaoPredador = Direcao(i)
			qtdePredadores++
			p.alterarIntensidadeEmocao(2, 0, 3)
			p.alterarQualidadeEmocao(-2, -3, 3)
		} else if campo.Agente == C_Presa {
			qtdePresasLivres++
		} else if campo.Agente == C_Presa_Fugindo {
			direcaoPresaFugindo = Direcao(i)
			qtdePresasFugindo++
			p.alterarIntensidadeEmocao(1, 0, 3)
			p.alterarQualidadeEmocao(-1, -3, 3)
		}
	}

	if qtdePredadores >= 3 {
		return p.morrer()
	} else if /*p.qualidadeEmocao < 0*/ qtdePredadores > 0 || qtdePresasFugindo > 0 {
		p.qtdeIteracaoPanico = 0
		var direcaoFuga Direcao
		if direcaoPredador > -1 {
			direcaoFuga = Direcao(direcaoPredador)
		} else {
			direcaoFuga = Direcao(direcaoPresaFugindo)
		}
		return p.fugir(direcaoFuga)
	} else {
		p.qtdeIteracaoPanico++
		p.alterarQualidadeEmocao(qtdePresasLivres, -3, 3)
		if p.qtdeIteracaoPanico > QtdeIteracaoPanico && qtdePresasLivres == 0 {
			p.alterarQualidadeEmocao(1, -3, 1)
			p.alterarIntensidadeEmocao(-1, 0, 3)
		}
		if p.qualidadeEmocao > 0 {
			p.mudouCor = false
		}
		return p.viver()
	}
}

func (p *Presa) fugir(direcao Direcao) (Posicao, PosMovimento) {
	p.mudouCor = true
	// vai na direcao oposta
	return p.moveAgente(ObterDirecaoOposta(direcao), 1)
}

func (p *Presa) morrer() (Posicao, PosMovimento) {
	p.morreu = true
	return Posicao{}, PosMovimento(-1)
}

func (p *Presa) getMorreu() bool {
	return p.morreu
}
