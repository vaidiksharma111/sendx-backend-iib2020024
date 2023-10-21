async function crawlURL() {
    const url = document.getElementById('url').value;
    const isPaying = document.getElementById('isPaying').checked;
    const response = await fetch(`/crawl?url=${url}&isPaying=${isPaying}`, {
        method: 'POST',
        body: JSON.stringify({ url: url }),
        headers: { 'Content-Type': 'application/json' }
    });
    const data = await response.text();
    document.getElementById('output').innerHTML = escapeHtml(data).replace(/\n/g, '<br>');
}

function escapeHtml(unsafe) {
    return unsafe
        .replace(/&/g, "&amp;")
        .replace(/</g, "&lt;")
        .replace(/>/g, "&gt;")
        .replace(/"/g, "&quot;")
        .replace(/'/g, "&#039;");
}

function Back() {
    window.location.href = `/`;
}