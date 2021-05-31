var myHeaders = new Headers();

var myInit = { 
    method: 'GET',
    headers: myHeaders,
    mode: 'cors',
    cache: 'default' 
};

var myRequest = new Request('phones', myInit);

fetch(myRequest).then(function (response) {
    var contentType = response.headers.get("content-type");
    if (contentType && contentType.indexOf("application/json") !== -1) {
        return response.json().then(function (data) {
            data.forEach(item => {
                insertRowData(item)
            })
        });
    } else {
        alert("Oops, we haven't got JSON!")
        console.log("Oops, we haven't got JSON!");
    }
});

function insertRowData(rowData){
    var tbodyRef = document.getElementById('listPhonesTable').getElementsByTagName('tbody')[0];
    var newRow = tbodyRef.insertRow();

    newRow.insertCell().appendChild(document.createTextNode(rowData['Country']));
    newRow.insertCell().appendChild(document.createTextNode(rowData['State']));
    newRow.insertCell().appendChild(document.createTextNode(rowData['CountryCode']));
    newRow.insertCell().appendChild(document.createTextNode(rowData['PhoneNumber']));
}