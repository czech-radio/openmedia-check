# Zadání

## Motivace a cíle

Pro účely analytiky vysílání se vyhotovují exporty dat ze systému OpenMedia v podobě XML souborů. Tyto soubory se každý den ukládají na sdílený disk do adresáře `$env:annova`. Tyto exportované soubory jsou dvojího druhu tkzv. rundowny, prefixované jako `RD_` a kontakty, prefixované jako `CT_`. Pro další zpracování třídíme tyto soubory automaticky do adresářů pomocí konzolových skriptů/programů [*openmedia*](https://github.com/czech-radio/openmedia/). Soubory představující rundowny se přesouvají do adresáře `$env:annova/Rundowns` a kontakty do adresáře `$env:annova/Contacts`. V těchto adresářích se dále třídí podle roku a týdne ke kterému přísluší např. 

```shell
Rundowns/
  2022/
    W01
    W02
    ...
```

V případě souborů představujích rundowny se export provádí hromadně jednou za den. Naproti tomu, export souborů představujících kontakty, se provádí kontinuálně po celý den, tak jak se záznamy upravují v aplikaci (systému) OpenMedia. Třídění je zcela v režii zmíněného programu *openmedia*, který pro účely třídění pracoval s datumem úpravy daného souboru. Soubor má sice datum zapsán i v názvu souboru, formát se však občas u různých souborů lišil, proto se přednostně použilo datum úpravy souboru. To se časem ukázalo jako možný **problém**. Pokud z nějakého důvodu nesouhlasí datum modifikace s datumem uvedeným v názvu souboru, pak je soubor špatně zařazen. Toto se stalo např. pokud byl výpadek systému a exporty za několik dní byly do adresáře přidány v jeden den. Dále se může soubor omylem upravit v nějakém programu, což také změní datum úpravy souboru.

**Cílem** vznikajícího programu je informovat a zamezit špatnému zařazení rundownů a kontaktů tím, že se automaticky a pravidelné kontroluje obsah daných adresářů na disku `$env:annova`.

## Požadované funkce

- [ ] Program má konzolové rozhraní. 
- [ ] Vstupem jsou cesty k vyšetřovaným adresářům. 
- [ ] Výstupem jsou informace o případných chybách a souhrnná statistika o obsahu daných adresářů v podobě JSON objektu.

### Kontrola chybovosti

Program pro zadané adresáře zkontroluje, jestli v nich umístěné OpenMedia XML soubory, náleží do adresáře, ve kterém je právě umístěný. Soubor má v názvu uvedené datum tzn., že kontolujeme, jetli uvedené datum v názvu souboru odpovídá adresáři. Soubory řadíme podle let a týdnů např. 

### Souhrnná statistika

Program by měl podat informaci o množství souborů ve vyšetřovaných adresářích.


