# Zadání

## Úvod

Pro účely analytiky vysílání se vyhotovují exporty dat ze systému OpenMedia v podobě XML souborů. Tyto soubory se každý den ukládají na sdílený disk do adresáře `$env:annova`. Tyto exportované soubory jsou dvojího druhu tkzv. rundowny, prefixované jako `RD_` a kontakty, prefixované jako `CT_`. Pro další zpracování třídíme tyto soubory automaticky do adresářů pomocí konzolových skriptů/programů [*openmedia*](https://github.com/czech-radio/openmedia/). Soubory představující rundowny se přesouvají do adresáře `$env:annova/Rundowns` a kontakty do adresáře `$env:annova/Contacts`. V těchto adresářích se dále třídí podle roku a týdne ke kterému přísluší např. `$env:annova/Rundowns/2022/W01` nebo `$env:annova/Rundowns/2022/W21` atd.

V případě souborů představujích rundowny se export provádí hromadně jednou za den. Naproti tomu, export souborů představujících kontakty, se provádí kontinuálně po celý den, tak jak se záznamy upravují v aplikaci (systému) OpenMedia. Třídění je zcela v režii zmíněného programu *openmedia*, který pro účely třídění pracoval s datumem úpravy daného souboru. Soubor má sice datum zapsán i v názvu souboru, formát se však občas u různých souborů lišil, proto se přednostně použilo datum úpravy souboru. To se časem ukázalo jako možný **problém**. Pokud z nějakého důvodu nesouhlasí datum modifikace s datumem uvedeným v názvu souboru, pak je soubor špatně zařazen. Toto se stalo např. pokud byl výpadek systému a exporty za několik dní byly do adresáře přidány v jeden den. Dále se může soubor omylem upravit v nějakém programu, což také změní datum úpravy souboru.

## Motivace a cíle

**Cílem** vznikajícího programu je informovat a zamezit špatnému zařazení rundownů a kontaktů tím, že se automaticky a pravidelné kontroluje obsah daných adresářů na disku `$env:annova`.

## Požadované funkce

- [x] Program má konzolové rozhraní.
- [ ] Program je dobře dokumentován a otestován.
- [x] Program může být použit jako součást Unix pipeline tzn. správně pracuje s `stdin/stdout/stderr`.
- [x] Vstupem jsou cesty k vyšetřovaným adresářům.
- [x] Výstupem je **kontrola chybovosti** a **souhrnná statistika** o obsahu daných adresářů v podobě JSON objektu.
  - [x] Chyby budou jasně definovány pomocí kódu (čísla/zkratky) a popisu.

Jako textový výstup požadujeme JSON objekt z důvodů dalšího zpracování např. pomocí programu [jq](https://stedolan.github.io/jq/).
Přesná podoba JSON objektu je zatím ponechána na zpracovateli a bude průbežně doplňována.

### Kontrola chybovosti

Program pro zadané adresáře zkontroluje, jestli v nich umístěné OpenMedia XML soubory, náleží do adresáře, ve kterém je právě umístěný. Soubor má v názvu uvedené datum tzn., že kontolujeme, jetli uvedené datum v názvu souboru odpovídá adresáři.

### Souhrnná statistika

Program by měl podat informaci o množství souborů ve vyšetřovaných adresářích (poměr správně zařazených souboru vs celkové množství souborů).


## Poznámky vývojářů

- Pokud máš v Go pole, pak jeho délku zjistíš pomocí `len()`
- Vracej hodnoty z funkcí, proč stále někde něco vypisuješ?
  - Nepoužíbvej Print, zakaž si to, vede tě to pak k netestování!

- Neotvírej a nezapisuj do souboru v jedné funkci.
  - To samé co s `fmt.Println()` . Obecně otevírat soubor, číst hodnoty a provádět výpočty v jedné funkci je netestovatelné. To je stále stejný problém, ať už programuješ v čemkoliv.

- Duplicity, *copy-paste* v několika  funkcích to samé jen změna jednoho textu nebo v jedno funkci třikrát `filepath.Join(folder, fn.Name())`.
