async function updateConfiguration() {
    const maxCrawlsValue = document.getElementById('maxCrawls').value;
    const numWorkersValue = document.getElementById('numWorkers').value;

    // Check if the input fields are not empty and are valid numbers
    if (!numWorkersValue && !maxCrawlsValue) {
        alert("Please enter valid numbers in both fields.");
        return;
    }

    await fetch('/config/numWorkers', {
        method: 'POST',
        body: numWorkersValue,
        headers: { 'Content-Type': 'application/json' }
    });

    await fetch('/config/maxCrawlsPerHour', {
        method: 'POST',
        body: maxCrawlsValue,
        headers: { 'Content-Type': 'application/json' }
    });

    // Clear the input fields
    document.getElementById('maxCrawls').value = '';
    document.getElementById('numWorkers').value = '';
}

async function retrieveConfiguration() {
    const response = await fetch('/config');
    const configData = await response.json();
    document.getElementById('configOutput').textContent = JSON.stringify(configData, null, 2);
}

function clearInputFields() {
    document.getElementById('maxCrawls').value = '';
    document.getElementById('numWorkers').value = '';
}

async function fetchCurrentConfiguration() {
    try {
        const response = await fetch('/get-config');
        if (!response.ok) {
            throw new Error('Network response was not successful');
        }
        const configData = await response.json();
        document.getElementById('numWorkers').value = configData.numWorkers;
        document.getElementById('maxCrawls').value = configData.maxCrawlsPerHour;
    } catch (error) {
        console.error('An error occurred during the fetch operation:', error.message);
    }
}

function navigateBack() {
    window.location.href = `/`;
}
