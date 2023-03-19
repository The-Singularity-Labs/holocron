# Holocron
[![Hippocratic License HL3-BDS-BOD-CL-ECO-EXTR-FFD-MEDIA-MY-SOC-SUP-SV-TAL-USTA-XUAR](https://img.shields.io/static/v1?label=Hippocratic%20License&message=HL3-BDS-BOD-CL-ECO-EXTR-FFD-MEDIA-MY-SOC-SUP-SV-TAL-USTA-XUAR&labelColor=5e2751&color=bc8c3d)](https://firstdonoharm.dev/version/3/0/bds-bod-cl-eco-extr-ffd-media-my-soc-sup-sv-tal-usta-xuar.html)
[![forthebadge](https://forthebadge.com/images/badges/contains-technical-debt.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/uses-badges.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/you-didnt-ask-for-this.svg)](https://forthebadge.com)


![](https://media1.giphy.com/media/ToMjGpsQduVVwLpZCxy/giphy.gif)

[Try it out](https://the-singularity-labs.github.io/holocron/index.html)

## About
This is a tool to build QR codes with gatekeepers, meant to embed & protect information in physical objects, called a Holocron. Holocrons are small ciphers that reveal a `treasure` when the right response to a prompt is provided. Holocron ciphers are small enough to be embedded in QR codes (and therefore [physical objects](https://qalo.com/collections/qr-dog-id-tags)) or NFC tags. 

The resulting Holocron should not be posted on the public internet. Brain key's aren't considered secure for [blockchain wallets](https://en.bitcoin.it/wiki/Brainwallet). However, in this case you can create a brain key (Holocron) that does not generate a wallet directly but rather points to a more secure wallet mnenomic phrase. On top of that, by embedding the Holocron in an object and keeping it off the public internet, we create a second factor an attacker has to breach in order to access the treasure. The attacker has to both know the password to the Holocron but must also physically possess the Holocron in order to steal the treasure.

# Quickstart

## Required Dependencies

- [Golang](https://go.dev//)
- [Tinygo](https://tinygo.org/)
- [Node](https://nodejs.org/en)
- [Yarn](https://yarnpkg.com/)



## Testing A Holocron 

![test](https://i.ibb.co/6X4L0RJ/test.png)

- Download or take a picture of the QR code with your phone
- Use a [ZXing](https://zxing.org/w/decode.jspx) based decoder ([like this one](https://online-barcode-reader.inliteresearch.com/)) to decode the 
- Paste the resulting data uri in your browser 
- Figure out the answer to the Gatekeeper prompt
- Use the `Decrypt` in this directory (or any Balloon Hashing implementation) function along with the cipher and answer to decrypt

## Creating a Holocron

### Asumptions

- The information you'd like to protect (`Treasure`) is relatively short (under 100 characters)
- The information does not need to be frequently accessed, this is better for cold storage
- The answer to your prompt is not trivially guessible 

### Generate Encrypted Key

- `make HOLOCRON_NAME=test HOLOCRON_GATEKEEPER="2+2?" HOLOCRON_ASCERTAINMENT=4 HOLOCRON_TREASURE=foobar"`
- Open .build folder to see QR code of holocron
- Copy this key to use in creating a Holocron

### Physical Storage

QR codes are a great mechanism for physical storage since they can be manifested physically in practically any way an image can. Options include:

- Printing on a [Metal Plate](https://bayphoto.com/)
- [Lamination](https://www.fedex.com/en-us/office/binding-laminating-finishing-services.html?cmp=KNC-10000002-0-0-0-FXO-US-US-EN-AISFXO121510430&gclid=Cj0KCQiA95aRBhCsARIsAC2xvfxyFgrJqhUobH4TRA4CIT3g1DxGe2nC575DHcMcY8M7K1ZqGhmgXh4aAjAzEALw_wcB&gclsrc=aw.ds)
- [Ceramic Tiling](https://www.zazzle.com/qr+code+tiles)
- [Archival Focused Paper](https://www.futurepkg.com/best-paper)
- [Optical Storage](https://www.amazon.com/Best-Sellers-External-CD-DVD-Drives/zgbs/pc/1292121011)

NFC tags have also been proven to be a practical means of storage. Tags are cheap and inconspicous.

Harddrives, SSDs, and flashdrives aren't great archival storage methods. Flash drives don't last that long and tape requires climate control to remain in good condition. A metal plate is small enough to fit in a safe or safety deposit box while also being able to survive adverse climates or even fires. 

### Types of Treasures to Store

- Blockchain Keys
- Master Key for Password Manager or 2 Factor Authenticator
- Coordinates to buried treasure if you have some
- Family Secrets
- Secret Recipes/Techniques
- Import Facts/Dates about your life
- etc

