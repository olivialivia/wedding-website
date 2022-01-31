window.addEventListener("load", function () {
    let form = document.getElementById("rsvp-form");
    let log = document.getElementById("submitted-form");
    let btn1 = document.getElementById("btn-confirm");

    let declined = document.getElementById("declined");
    declined.addEventListener("change", (event) => {
        let inputs = document.getElementsByClassName("accept-inputs");
        for (let i = 0; i < inputs.length; i++) {
            let input = inputs[i];
            input.setAttribute("disabled", "");
        }
    });

    let accepted = document.getElementById("accepted");
    accepted.addEventListener("change", (event) => {
        let inputs = document.getElementsByClassName("accept-inputs");
        for (let i = 0; i < inputs.length; i++) {
            let input = inputs[i];
            input.removeAttribute("disabled");
        }
    });

    btn1.addEventListener("click", async (event) => {
        event.preventDefault();
        const formData = new FormData(form);

        let success = false;

        try {
            let response = await fetch("https://dendrite.kegsay.com/_wedding/submit", {
                method: "POST",
                body: formData,
            });
            success = response.ok;
        } catch (err) {
            console.log("request error", err);
        }

        if (!success) {
            log.textContent =
                "Wystąpił problem. Spróbuj jeszcze raz lub skontaktuj się z nami poniżej.";
        } else {
            log.textContent = "Dziękujemy za wypełnienie!";
        }
    });
});
