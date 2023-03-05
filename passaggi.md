# Live-blog
Uso questo progetto per fare l'upload e il download di file criptati sul server invido.it.
Una specie di cloud primitiva. Inizialmente avevo pensato questo service per aggiornare il mio sito,
ma è una funzionalità della quale non ho mai sentito la mancanza.
Quella, invece, di inviare files da diversi dispositivi sul server e di scaricarli su altri
è una funzionalità della quale ho bisogno, specialmente in ambienti dove la chiavetta usb non funziona
più, ma esiste ancora la possibilità di eseguire uploads sul mio server.  

A grandi linee le funzionalità del server sono:
- Sul server le risorse (in genere files) sono criptate. La chiave privata è sul client.
- Download di una risorsa così com'è
- Upload di una risorsa così com'è. Se è lo stessa viene sovrascritta. 
- Lista delle risorse
- Cancellare una risorsa
- Interfaccia web (al momento non necessaria)

Per il Client vorrei avere un tool a linea di comando che esegue:
- Criptazione della risorsa con chiave privata del client
- Upload del file criptato
- Download del file criptato e decriptazione finale
- Lista delle risorse sul server
- Cancellare una risorsa

## API

list: Route API/list, validation: jwt in auth header. Mostra la lista delle risorse.

## Token e Refresh token
Ci sono due tipi di token, uno per il refresh (Token.RefreshToken) e uno per l'autenticazione (Token.AccessToken). 
Il token refresh viene usato con il metodo Token nella API senza usare lo username. 
Questo ritorna un nuovo Token completo (Token.AccessToken) che può essere usato nelle altre API. Il Refresh viene usato per non chiedere la password all'utente dopo la prima chiamata.
Così si ha un'autenticazione che si rinnova automaticamente. 
È possibile, nel client, l'uso dell'hash della password per generare il token (shared secret).


## Compilare per linux
Apri una nuova powershell e poi:

    $env:GOOS = "linux"
    go build -o live-blog.bin
Con WLC si può controllare che il live-blog.bin funziona correttamente.

## Aggiornare il service
- Crea una nuova versione (cambio in idl.go)
- Crea il file live-blog.bin per linux 
- Usa .\deploy -target invido
- In WLC lancia ./copy_app_to_invido.sh
- Su invido: ./update-service.sh

## Deployment di live.invido.it
Per prima cosa aggiornare il DNS aggiungendo live.invido.it
Poi bisogna abilitare il reverse proxy con https su nginx. Parto dal file iolvienna:
/etc/nginx/sites-available$ sudo cp vienna.invido.it  live.invido.it
poi cambio i link http e il nome del server usando i dati del nuovo service live-blog
sudo nano live.invido.it
Bisogna abilitare il sito:
sudo ln -s /etc/nginx/sites-available/live.invido.it  /etc/nginx/sites-enabled/live.invido.it
Un test della configurazione di nginx:
sudo nginx -t
Ora un restart:
sudo systemctl restart nginx

Su invido.it poi si aggiorna il certificato (NOTA se questo step si fa prima di nginx enable e available, certbot va a modificare il file default e 
il sito live.invido.it non funziona. Errore "nginx: [warn] conflicting server name "live.invido.it" on [::]:443, ignored". 
Qui bisogna modificare il file default cancellando il redirect del sito live su invido.it):
sudo certbot --nginx -d <lista di tutti i domini> -d live.invido.it
Alla scelta scegliere [1], non modificare nginx (anche se come messo sopra, certbot va a cambiare il file di default)

Ora va installata la app. Uso la dir:
~/app/go$mkdir live-blog\zips
copio lo zip deployed locale in live-blog\zips con ./copy_app_to_invido.sh
copio ./update-service.sh in ~/app/go/live-blog e lo lancio per scompattare lo zip nella dir ./current
Poi si va ./current e si prova il service con: ./live-blog.bin

Poi si mette il programma live-blog.bin come service di sistema.
sudo nano /lib/systemd/system/live-blog.service
Si mette tutto il necessario che non sto qui a riportare (readme_hetzner)
Abilitare il service:
sudo systemctl enable live-blog.service
Ora si fa partire il service (resistente al reboot):
sudo systemctl start live-blog
Per vedere i logs si usa:
sudo journalctl -f -u live-blog

## Usare il live service
Per prima cosa va creato un account, per esempio admin. Live supporta un solo account.
Esso va creato a liene di comando.  Per cambiare le credential bisogna cancellare il file json.

## Sign In
Uso il jwt token per il sign in. Non credo sia necessaria una session sul server.
Le richieste avvengono via rest con auth token.
Questo quello che si dovrebbe fare coi token:
3) The client (Front end) will store refresh token in his local storage and access token in cookies.

## Vuetify e Icon materialize
Il riferimento per le icons: https://materializecss.com/icons.html
Il riferimento per i componenti: https://vuetifyjs.com/en/components/text-fields/


## Chiave pubblica
Per validare il token jwt occorre la chiave pubblica
openssl rsa -in key.pem -pubout -out pubkey.pem

