package db

import (
	"context"
	"log"

	"github.com/orbit-alliance/orbit-backend/internal/domain/coin"
	"github.com/orbit-alliance/orbit-backend/internal/domain/nft"
)

func SeedGoodActions(ctx context.Context, repo *GoodActionRepository) {

	// Verifica se já existem GoodActions na coleção
	existingActions := repo.LoadAll(ctx)
	if len(existingActions) > 0 {
		log.Println("GoodActions já existem na coleção. Pulando o seed.")
		return
	}

	actions := []*coin.GoodAction{
		coin.NewGoodAction(
			"Frequência Premiada",
			`A ação Frequência Premiada tem como objetivo incentivar os cadetes a frequentarem regularmente o campus da 42 Rio, reforçando a cultura de aprendizado entre pares e promovendo a troca de conhecimento. O engajamento presencial melhora a colaboração, facilita o avanço nos projetos e fortalece a comunidade acadêmica.

Organização
Funcionamento: o cadete recebe uma bonificação diária apenas por estar presente no campus e conectado à intra por um tempo mínimo. Para aumentar o incentivo à regularidade, há um sistema de multiplicador de frequência (streak), que oferece bônus adicionais para aqueles que mantêm uma sequência ininterrupta de presença.

Critérios
- Para receber a recompensa diária, o cadete deve logar na intra.
- O sistema identificará automaticamente a frequência e registrará o saldo de recompensas na conta do cadete.
- O multiplicador de frequência (streak) será ativado para cadetes que mantiverem presença contínua no campus.
- Se o cadete faltar um único dia, o streak é perdido e reinicia do zero.`,
			1,
			nil,
			[]nft.NFT{},
			[]nft.NFT{},
		),
		coin.NewGoodAction(
			"Bônus de Projeto",
			`Objetivo
O sistema Bônus de Projetos tem como objetivo incentivar os cadetes a realizarem os desafios extras dos projetos, explorando conteúdos mais avançados e aprimorando suas habilidades técnicas. Atualmente, muitos cadetes evitam os bônus devido à complexidade ou ao tempo necessário para concluí-los. Com este incentivo, os bônus se tornam oportunidades valiosas de aprendizado, além de uma forma de acumular Galactos.

Organização
Cada projeto na 42 possui uma estrutura básica (mandatória) e desafios adicionais, os bônus. Para incentivar os cadetes a completarem essas partes extras, será concedida uma bonificação em Galactos proporcional ao nível do bônus concluído.
A recompensa é ajustada de acordo com a dificuldade e complexidade do bônus dentro do projeto, podendo variar de 1₲ até 25₲ por projeto.

Critérios
- O cadete precisa entregar e ter validado o bônus do projeto na intra.
- Apenas bônus reconhecidos pela intra serão elegíveis para bonificação.
- A quantidade de Galactos concedidos varia de acordo com o projeto e sua pontuação.
- O pagamento dos Galactos será automático, assim que a intra registrar a conclusão do bônus.`,
			25,
			nil,
			[]nft.NFT{},
			[]nft.NFT{},
		),
		coin.NewGoodAction(
			"Participação em Talk",
			`Objetivo
A ação Participação em Talks tem como objetivo incentivar os cadetes a comparecerem às palestras e eventos acadêmicos promovidos dentro da 42 Rio. 

Talks são momentos de troca de conhecimento e networking, fundamentais para o crescimento técnico e profissional dos cadetes. A presença em eventos como esses amplia a visão do cadete sobre tecnologia, inovação e carreira, além de contribuir para a cultura colaborativa da escola.

Organização
Para garantir a participação ativa nos Talks, será implementado um sistema de registro de presença.
A validação será feita de forma automática, garantindo que apenas cadetes presentes no Talk recebam a recompensa.

Importante: Para que os Galactos sejam creditados, o evento precisa ter um quórum mínimo de 10 cadetes. Caso esse número não seja atingido, a participação não será bonificada.

Critérios
- O sistema será programado para validar automaticamente a quantidade mínima de participantes antes de distribuir as recompensas.
- Presença obrigatória no evento até o final.
- O evento deve contar com pelo menos 10 cadetes presentes para gerar bonificação.

Recompensa
15₲ por participação validada em um Talk.`,
			15,
			nil,
			[]nft.NFT{},
			[]nft.NFT{},
		),
		coin.NewGoodAction(
			"26 dias Consecutivos de Presença",
			`Objetivo
A ação 26 dias consecutivos de presença tem como objetivo incentivar os cadetes a manterem uma rotina de frequência regular no campus da 42 Rio. A presença contínua é fundamental para o aprendizado colaborativo e para o fortalecimento da comunidade acadêmica.
Organização
Para garantir a participação ativa dos cadetes, será implementado um sistema de registro de presença. A validação será feita de forma automática, garantindo que apenas cadetes presentes no campus recebam a recompensa.
Critérios
- O sistema será programado para validar automaticamente a frequência dos cadetes.
- Presença obrigatória no campus por 26 dias consecutivos.
- O sistema de multiplicador de frequência (streak) será ativado para cadetes que mantiverem presença contínua no campus.
- Se o cadete faltar um único dia, o streak é perdido e reinicia do zero.
Recompensa
- 100₲ por manter 26 dias consecutivos de presença no campus.`,
			100,
			nil,
			[]nft.NFT{},
			[]nft.NFT{},
		),
		coin.NewGoodAction(
			"42 dias Consecutivos de Presença",
			`Objetivo
A ação 42 dias consecutivos de presença tem como objetivo incentivar os cadetes a manterem uma rotina de frequência regular no campus da 42 Rio. A presença contínua é fundamental para o aprendizado colaborativo e para o fortalecimento da comunidade acadêmica.
Organização
Para garantir a participação ativa dos cadetes, será implementado um sistema de registro de presença. A validação será feita de forma automática, garantindo que apenas cadetes presentes no campus recebam a recompensa.
Critérios
- O sistema será programado para validar automaticamente a frequência dos cadetes.
- Presença obrigatória no campus por 42 dias consecutivos.
- O sistema de multiplicador de frequência (streak) será ativado para cadetes que mantiverem presença contínua no campus.
- Se o cadete faltar um único dia, o streak é perdido e reinicia do zero.
Recompensa
- 150₲ por manter 42 dias consecutivos de presença no campus.`,
			150,
			nil,
			[]nft.NFT{},
			[]nft.NFT{},
		),
	}

	for _, action := range actions {
		err := repo.Save(ctx, action)
		if err != nil {
			log.Printf("Erro ao salvar GoodAction %s: %v", action.Name, err)
		} else {
			log.Printf("Seeded GoodAction: %s", action.Name)
		}
	}
}
