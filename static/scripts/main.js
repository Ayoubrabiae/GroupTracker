const bodyElement = document.body

// Deal with local storage
const localStorageToggle = (key) => {
    if (localStorage.getItem(key)) {
        localStorage.setItem(key, "")
    } else {
        localStorage.setItem(key, "true")
    }
}

// Dark mode functionality
const darkModeBtn = document.querySelector(".dark-mode-btn")
if (localStorage.getItem("dark")) {
    bodyElement.classList.add("dark")
}

if (darkModeBtn) {
    if (localStorage.getItem("dark")) {
        darkModeBtn.children[0].classList.add("bxs-sun")
    }
    
    darkModeBtn.addEventListener("click", () => {
        localStorageToggle("dark")
        bodyElement.classList.toggle("dark")
        darkModeBtn.children[0].classList.toggle("bxs-sun")
    })
}

// Hero functionality
const heroImageElements = document.querySelectorAll(".hero .images img")

if (heroImageElements.length) {
    const artistImages = ["queen.jpeg", "soja.jpeg", "pinkfloyd.jpeg", "scorpions.jpeg",
        "xxxtentacion.jpeg", "macmiller.jpeg", "joynerlucas.jpeg", "kencricklamar.jpeg",
        "acdc.jpeg", "pearljam.jpeg", "katyperry.jpeg", "rihanna.jpeg", "genesis.jpeg",
        "philcollins.jpeg", "ledzeppelin.jpeg", "jimihendrix.jpeg", "beegees.jpeg", "deeppurple.jpeg",
        "aerosmith.jpeg", "direstraits.jpeg", "30.jpeg", "imagineDragons.jpeg", "juiceWrld.jpeg", "logic.jpeg",
        "alecBeenjamin.jpeg", "bobbyMcFerrin.jpeg", "R3HAB.jpeg", "postmalone.jpeg", "travisScott.jpeg", "jCole.jpeg",
        "nickelback.jpeg", "mobbDeep.jpeg", "gunsNRoses.jpeg", "nwa.jpeg", "u2.jpeg", "arcticmonkeys.jpeg", "fallOutBoy.jpeg",
        "gorillaz.jpeg", "eagles.jpeg", "linkinpark.jpeg", "redhotchilipeppers.jpeg", "eminem.jpeg", "greenday.jpeg", "metallica.jpeg",
        "coldplay.jpeg", "maroon5.jpeg", "twentyonepilots.jpeg", "therollingstones.jpeg", "muse.jpeg", "foofighters.jpeg", "thechainsmokers.jpeg"]

    setInterval(() => {
        const randomElementNum = Math.floor(Math.random() * heroImageElements.length)
        const randomImageNum =  Math.floor(Math.random() * artistImages.length)
        const imageSrc = `https://groupietrackers.herokuapp.com/api/images/${artistImages[randomImageNum]}`
        heroImageElements[randomElementNum].src = imageSrc
    }, 3000)
}

// Top Button Functionality
const topBtn = document.querySelector(".top-arrow")

if (topBtn) {
    topBtn.addEventListener("click", () => {
        window.scrollTo(0, 0)
    })

    window.addEventListener("scroll", () => {
        if (window.scrollY > 400) {
            if (!topBtn.classList.contains("show")) {
                topBtn.classList.add("show")
            }
        } else {
            if (topBtn.classList.contains("show")) {
                topBtn.classList.remove("show")
            }
        }
    })

}

/*Artist Details Slide Cards*/
let items = document.querySelectorAll('.info .info-child');
let prev = document.getElementById('prev');
let next = document.getElementById('next');

let active = 3;
function loadShow(){
    let setting = 0;
    items[active].style.transform = `none`;
    items[active].style.zIndex = 1;
    items[active].style.filter = 'none';
    items[active].style.opacity = 1;
    for(var i = active + 1; i < items.length; i++){
        setting++;
/*${1 - 0.2*setting} is a dynamic value computed using the setting variable. If setting is 0, the scale factor is 1 (no scaling). For example, if setting is 3,
the scale factor becomes 1 - 0.2*3 = 1 - 0.6 = 0.4. This would scale the element down to 40% of its original size.*/
        items[i].style.transform = `translateX(${120*setting}px) scale(${1 - 0.2*setting}) perspective(16px) rotateY(-1deg)`;
        items[i].style.zIndex = -setting;
        items[i].style.filter = 'blur(5px)';
        items[i].style.opacity = setting > 2 ? 0 : 0.6;
    }
    setting = 0;
      for(var i = active - 1; i >= 0; i--) {
        setting++;
        items[i].style.transform = `translateX(${-120*setting}px) scale(${1 - 0.2*setting}) perspective(16px) rotateY(1deg)`;
        items[i].style.zIndex = -setting;
        items[i].style.filter= 'blur(5px)';
        items[i].style.opacity = setting > 2 ? 0 : 0.6;
    }
}

if (items.length) {
    loadShow();
    next.onclick = function() {
        /*if its the last elemnt we add 1 else we we input it without anything*/
        active = active + 1 < items.length ? active + 1 : 0;
        loadShow();
    }
    
    prev.onclick = function() {
        active = active - 1 >= 0 ? active - 1 : items.length -1;
        loadShow();
    }
}

// exit button
const cardsExitBtn = document.querySelector(".profile #exit")

if (cardsExitBtn) {
    cardsExitBtn.addEventListener("click", () => {
        if (document.referrer != window.location.origin+"/") {
           window.location.assign(window.location.origin)
            return
        }
        window.history.back()
    })
}