function fetchData(){
    var myHeaders = new Headers();
    var myInit = { 
        method: 'GET',
        headers: myHeaders,
        mode: 'cors',
        cache: 'default'
    };
    var filters = {
        country: document.getElementById('filterCountry').value,
        statePhone: document.getElementById('filterStatePhone').value
    }
    var urlFetchData = 'phones?country=$country&state=$state'.replace('$country', filters.country).replace('$state', filters.statePhone)
    var myRequest = new Request(urlFetchData, myInit);

    fetch(myRequest).then(function (response) {
        var contentType = response.headers.get("content-type");
        if (contentType && contentType.indexOf("application/json") !== -1) {
            return response.json().then(function (data) {
                data.forEach(item => {
                    insertRowData(item)
                })
            });
        } else {
            console.log("JSON FAIL!");
        }
    });
}

function insertRowData(rowData){
    var tbodyRef = document.getElementById('listPhonesTable').getElementsByTagName('tbody')[0];
    var newRow = tbodyRef.insertRow();

    newRow.insertCell().appendChild(document.createTextNode(rowData['Country']));
    newRow.insertCell().appendChild(document.createTextNode(rowData['State']));
    newRow.insertCell().appendChild(document.createTextNode(rowData['CountryCode']));
    newRow.insertCell().appendChild(document.createTextNode(rowData['PhoneNumber']));
}

function updateTablePhone(){
    let tbodyPhoneTable = document.getElementById("listPhonesTable").querySelectorAll('tbody')[0]
    for (let index = 1; index < tbodyPhoneTable.rows.length; index++) {
        tbodyPhoneTable.deleteRow(0)
    }
    fetchData()
}

fetchData()