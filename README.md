# Coding Task 2
Rahmenbedingungen:
- Bitte die Aufgabe selbstständig und allein lösen, natürlich darf dabei Literatur und / oder das Internet benutzt werden.
- Bitte teilen Sie uns Ihre Bearbeitungszeit mit.
- Die Lösung hätten wir gern als GitHub Projekt.
- Die Programmiersprache ist frei wählbar, wobei eine ‘gewöhnliche’ Sprache bevorzugt wäre.
- Das Resultat muss von uns gebaut und ausgeführt werden können. Bitte entsprechende build-scripte oder Makefiles bereitstellen.
- Eigene Annahmen und wichtige Implementierungsdetails bitte klar kommentieren.
- Eventuelle sinnvolle Zwischenergebnisse dürfen gern als separater Git commit vorliegen. 
- Genauso wichtig wie das lauffähige Programm ist die Dokumentation (readme und code comments) der Lösungsidee und der einzelnen Programmteile und Tests.

**Das Hauptziel ist es das wir erleben wie Sie Software in einem professionellen Umfeld entwickeln. Die gesamte Bearbeitungsdauer sollte max. 1-2 Tage sein.**

## Aufgabe:

Implementieren Sie eine Funktion **MERGE** die eine Liste von Intervallen entgegennimmt und als Ergebnis wiederum eine Liste von Intervallen zurückgibt. Im Ergebnis sollen alle sich überlappenden Intervalle gemerged sein. Alle nicht überlappenden Intervalle bleiben unberührt.

**Beispiel:**
Input: [25,30] [2,19] [14, 23] [4,8]  Output: [2,23] [25,30]

Wie ist die Laufzeit Ihres Programms ?  
Wie kann die Robustheit sichergestellt werden, vor allem auch mit Hinblick auf sehr große
Eingaben ?  
Wie verhält sich der Speicherverbrauch ihres Programms ?

# Lösung:
Programmiersprache: Golang
Entwicklungsdauer: 2 Stunden 30 Minuten
## Ordnerstruktur
```
├── Makefile
├── README.md
├── main.go
└── merg
    ├── merge.go
    └── merge_test.go

1 directory, 5 files

```
Funktion befindet sich im ```merge/merge.go```.  
Test im ````merge/merge.go````

## Verwendung des Makefiles
Kompiliert die Anwendung
```
make build
```
Führt die Anwendung aus
```
make run
```
Kompiliert die Anwendung für alle Betriebssysteme
```
make compile
```


## Annahmen:
- Die erste Zahl in einem interval ist immer die kleinste
- Zahlen sind Integer
- Alle Zahlen sind Positiv

## Vorgehensweise
1. Aufgabenstellung untersuchen Anforderungen festmachen
2. Eingabe- und Rückgabeparameter festlegen (definieren der Typen in Go)
3. Tests erstellen für die Anwendung
4. Eingabe sortieren (Hilfsfunktion erstellen)
5. Neue liste erstellen mit erstem sortierten wert als Startintervall
6. Über die Liste iterieren und die einzelnen Intervalle mit den bereits gespeicherten Vergleichen
7. Hilfsfunktion erstellen um den größten Wert zurückzuliefern
8. Fertige Liste zurückgeben

## Laufzeit
### _time_ Befehl Auswertung
```
$ time go run main.go
$ [[2 23] [25 30]]
$ go run main.go  0.20s user 0.14s system 142% cpu 0.244 total
```
### go Benchmark
```
$ go test -bench .
$ goos: darwin
$ goarch: amd64
$ pkg: gitlab.com/skywalker_etw/DaimlerTssMerge/merge
$ BenchmarkMerge-16        6743065               165 ns/op
$ PASS
$ ok      gitlab.com/skywalker_etw/DaimlerTssMerge/merge  1.347s

```
## Robustheit
Die Robustheit der Anwendung kann sichergestellt werden, indem die Eingaben auf ihre zugelassenen Werte überprüft werden.
In diesem Fall wird dies durch die Verwendung einer typensicheren Programmiersprache implementiert.
Da Go mit vordefinierten Typen arbeitet und die Anwendung nur mit den vorgesehenen Datentypen funktionsfähig ist.
## Speicherverbrauch
