const bodyElement = document.body
const headerElement = document.querySelector("header")
const darkModeBtn = document.querySelector(".dark-mode-btn")

darkModeBtn.addEventListener("click", () => {
    console.log("Welcome")
    bodyElement.classList.toggle("dark")
    darkModeBtn.children[0].classList.toggle("bxs-sun")
})