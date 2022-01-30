window.addEventListener("load", function () {
    let form = document.getElementById("rsvp-form");
    let log = document.getElementById("submitted-form");
    let btn1 = document.getElementById("btn-confirm");

    btn1.addEventListener("click", async (event) => {
        log.textContent = "Dziękujemy za wypełnienie!";
        event.preventDefault();
        const formData = new FormData(form);

        let success = false;

        try {
            let response = await fetch("http://localhost:1906/_wedding/submit", {
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
        }
    });
});
