// main nav redirect on click and color the header if no highlight is visible
(function () {
    let highlighted = false;

    let items = document.querySelector('#main-nav').querySelectorAll('div');
    items.forEach((ele) => {
        // if there is one highlighted -> don't highlight the header
        ele.parentNode.classList.forEach((c) => {
            if (c === 'current-page') {
                highlighted = true;
            }
        });

        // add event listener to all buttons
        ele.parentNode.addEventListener('click', () => {
            window.location.href = ele.dataset.path;
        })
    });

    if (! highlighted) {
        const highlightColor = getComputedStyle(document.body).getPropertyValue('--main-hover-color');
        const header = document.querySelector('#header-title').querySelector('a');
        header.style.color = highlightColor;
    }
})();

// show the vertical nav bar when click action on small-nav-header (when screen is small - mobile)
(function () {
    let nav = document.querySelector('#small-nav-header');
    nav.addEventListener('click', () => {
        const s = document.querySelector('#main-nav').style;
        s.display = s.display === 'block' ? 'none' : 'block';
    });
})();
