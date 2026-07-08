
// TODO Maybe more like this
// <span class="scramble" title="or good or bad or awful" alt="???">ok</span>

// TODO later words less likely?
// TODO pad words?
// TODO sort by word len and fill with glitches?
// TODO glitch in between words?
const hashmap = {}
function scramble (interval = 1000 * 32) {
    setTimeout(() => {
        const targets = document.getElementsByClassName('scramble');
        for (target of targets) {
            const text = target.innerHTML;
            const isFirstFetch = text.match(/^\{.*\}$/);
            let words = [];

            // FIXME Make much, much nicer
            if (isFirstFetch) {
                const key = isFirstFetch.input;

                target.setAttribute('title', key)

                words = Array.from(text.match(/\w+/g));
                hashmap[key] = words;
            } else {
                words = hashmap[target.title];
            }

            const word = words[Math.floor(Math.random() * words.length)];

            target.innerHTML = word
        }

        scramble();

    }, interval);
}

document.addEventListener("DOMContentLoaded", scramble);

