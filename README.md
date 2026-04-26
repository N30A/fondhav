# Fondhav

Fondhav är ett REST API byggt i Go för att söka och filtrera svenska fonders innehav. Data hämtas från [Finansinspektionens fondregister](https://www.fi.se/sv/vara-register/fondinnehav/). Fokus har lagts på Go:s standardbibliotek där externa beroenden endast används om det inte räcker till.

## Bakgrund

Vill man kolla fonders innehav är det oftast begränsat till topp 10 eller liknande. Alternativet är att ladda ned fondinnehaven från respektive fondleverantörs sida (om de ens publicerar detta). Finansinspektionen kräver att svenskregistrerade fonder redovisar sina innehav, vilket FI publicerar i XML-format. Detta kräver dock att man laddar ned och går igenom datan själv utan ett lättanvänt sökgränssnitt. Fondhav löser det genom att ladda ned, normalisera och exponera datan via ett API.

## Bygg

```bash
make build
```

## Tester

```bash
make test
```

## Licens

[MIT](LICENCE)
