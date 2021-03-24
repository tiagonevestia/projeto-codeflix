<h1 align="center"> Microsserviços WINEHOUSE </h1>

## Descrição do Projeto 👾
Converter vídeos no formato MP4 em um formato mais adequado para o playback de vídeos na internet: MPEG-DASH.

## Fluxo 🔄
* Recebemos uma mensagem via RabbitMQ informando qual o vídeo deve ser convertido.
* Fazemos o download do vídeo no Google Cloud Storage
* Fragmentamos o vídeo
* Convertemos o vídeo para MPEG-DASH
* Faz o upload do vídeo no Google Cloud Storage
* Envia uma notificação via fila com as informações do vídeo convertido ou informando erro na conversão.
* Em caso de erro, a mensagem original enviada via RabbitMQ será rejeitada e encaminhada diretamente a uma Dead Letter Exchange.

## Pontos Importantes ✅
* Sistema deverá processar diversas mensagens de forma paralela/concorrente.
* Um simples MP4 quando convertido para MPEG-DASH é segmentado em múltiplos arquivos de áudio e vídeo, logo o processo de upload não é apenas de um único arquivo.
* O processo de upload também deve ser concorrente.



