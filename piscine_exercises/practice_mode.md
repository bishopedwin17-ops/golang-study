# 🔥 Go Piscine — Practice Mode (Exam Simulator)

Use this document to simulate real checkpoint exam conditions.

---

## ⚙️ Setup

1. Open a **new blank file** (e.g. `scratch.go`)
2. Start a timer — see the time limits below
3. Solve the assigned exercises **from memory only** — no notes, no answers folder
4. When the timer ends, compare with `checkpoint02answers/`
5. Note what you missed in `common_mistakes.md`

---

## 🎰 How the Real Exam Works

- You are given **3–5 random exercises** from the pool
- You have roughly **3 hours**
- Only one exercise is delivered at a time — you submit, get the next
- The checker is byte-exact — spacing, newlines, output format all matter

---

## 🏋️ Practice Sessions

### Session Type A — "Speed Drill" (10 min per exercise)

Pick 3 exercises randomly from this list and solve each in 10 minutes:

```
countalpha  atoi       itoa       itoabase   checknumber
findprevprime  cameltosnakecase  wordflip   cleanstr   expandstr
rostring    inter      union      addprimesum  fprime
wdmatch     hiddenp    brackets   rpncalc    canjump
```

**How to pick randomly:** look at the clock — use the seconds digit × 4 as an index.

---

### Session Type B — "Full Mock Exam" (45 min total)

Pick exactly **4 exercises** — one from each tier:

| Tier | Pick one |
|---|---|
| Easy (functions) | `countalpha`, `checknumber`, `itoa`, `fromto` |
| Medium (string programs) | `cleanstr`, `inter`, `union`, `wdmatch` |
| Hard (programs) | `addprimesum`, `fprime`, `brackets`, `rpncalc` |
| Bonus | `romannumbers`, `findpairs`, `printrevcomb` |

Set a 45-minute timer. Stop when it rings, even if not done.

---

### Session Type C — "Weak Spots Only" (20 min)

Look at your `PROGRESS.md`. Find 3 exercises still marked `[ ]`.
Solve those specifically.

---

## ✅ After Each Session — Checklist

- [ ] Did I handle the `len(os.Args)` edge case?
- [ ] Did every output end with exactly `\n`?
- [ ] Did I avoid using the standard library where forbidden?
- [ ] Could I explain every line to someone else?

---

## 🧠 What to Do When You're Stuck

1. **Don't look at the answer immediately.** Give it 5 more minutes.
2. Break the problem into smaller steps on paper.
3. Write the expected output first, then work backwards.
4. If still stuck after 10 min → look at the answer, understand it line by line, close it, rewrite it from memory.

---

## 🚩 Common Traps on the Real Exam

| Trap | Why it fails | Fix |
|---|---|---|
| Using `fmt.Println` in z01 exercises | Outputs extra data | Use `z01.PrintRune` only |
| Forgetting the trailing `\n` | Byte-exact check fails | Always end with `\n` |
| Using `strings.Split` | Standard library not allowed | Split manually with a loop |
| Off-by-one in loops | Wrong output count | Test with boundary values |
| Not checking `len(os.Args)` | Panic on wrong input | Always check first |
| Using `strconv.Atoi` | Standard library not allowed | Write your own `Atoi` |
