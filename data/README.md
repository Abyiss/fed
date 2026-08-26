# Bundled routing data

The bundled FedACH and Fedwire directories are historical snapshots. They are not
a complete or continuously updated replacement for the Federal Reserve Banks'
E-Payments Routing Directory. For current production use, supply licensed current
directory files through the service's existing data-path or download configuration.

## Supplemental participant records

### Routing number 091311229

The FedACH and Fedwire records for Mercury's routing number at Choice Financial
Group were added from publicly available participant details:

- The Federal Reserve Banks' E-Payments Routing Directory identified routing
  `091311229` as a current FedACH participant for `MERCURY, PARTNERING WITH CHOICE
  BANK` and a Fedwire participant for `CHOICE FINANCIAL GROUP`.
- Mercury's published domestic transfer instructions corroborated the routing
  number, Choice Financial Group relationship, Fargo address, and support for ACH
  and domestic wires.

The source details were accessed on August 26, 2026:

- [Federal Reserve E-Payments Routing Directory](https://www.frbservices.org/resources/routing-number-directory/)
- [Mercury example domestic transfer instructions](https://cdn-development.mercury.com/demo-assets/example-wire-details.pdf)

The fixed-width records preserve the rail-specific names published by the Federal
Reserve. Where the public detail display omitted a fixed-width-only representation,
the record uses the format-defined current-view code (`1`), no ZIP extension
(`0000`), and no settlement-only designation (a blank field). Regression tests
assert these representations and the complete parsed participant details.
