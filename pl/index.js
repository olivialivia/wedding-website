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
            let response = await fetch("http://127.0.0.1:5501/submit", {
                method: "POST",
                body: formData,
            });
            success = response.ok;
        } catch (err) {
            console.log("request error", err);
        }

        if (!success) {
            log.textContent = "There was a problem submitting the form. Please retry or email us.";
        }
    });
});
