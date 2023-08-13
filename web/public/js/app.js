'use strict';

const serverUrl = 'http://localhost:8080/api';

const listVehicles = async () => {
  await axios
    .post(`${serverUrl}/vehicles`)
    .then(function (response) {
      const vehicles = response.data;
      let li = `<tr><th>Marca</th><th>Modelo</th><th>Ano</th><th>Preço</th><th>Cmd</th></tr>`;
      vehicles.forEach((vehicle) => {
        //const data = JSON.stringify(vehicle);
        const data = btoa(JSON.stringify(vehicle));
        const button = `<button onclick="listLeads('${data}');">leads</button>`;
        li += `<tr>
          <td>${vehicle.brand}</td>
          <td>${vehicle.model} </td>
          <td>${vehicle.year}</td>
          <td>${vehicle.price}</td>
          <td>${button}</td>
        </tr>`;
      });
      document.getElementById('tab-vehicles').innerHTML = li;
    })
    .catch(function (error) {
      console.log(error);
    });
};

(function () {
  function init() {
    var router = new Router([
      new Route('list-vehicles', 'vehicle/list.html', true),
    ]);
  }
  init();
  document.addEventListener('DOMContentLoaded', listVehicles());
})();

const listLeads = async (vehicle) => {
  const parsed = JSON.parse(atob(vehicle));

  let div = `
    <div>Marca: ${parsed.brand}</div>
    <div>Model: ${parsed.model}</div>
    <div>Ano: ${parsed.year}</div>
    <div>Preço: ${parsed.price}</div>
    <br/>
    <div><center><h2>Leads</h2></center></div>
    <div><center><button onclick="showForm('lead-container','fname');">novo</button></center></div>
  `;
  document.getElementById('tab-vehicles').style.display = 'none';
  document.getElementById('vehicle-leads').innerHTML = div;

  const leads = [];
  leads.push({
    name: 'Lucas',
    email: 'lucas@gmail.com',
    phone: '11999999999',
  });

  let li = `<tr><th>Nome</th><th>E-mail</th><th>Telefone</th><th>Cmd</th></tr>`;
  leads.forEach((lead) => {
    const btnRemove = `<button>excluir</button>`;
    li += `<tr>
      <td>${lead.name}</td>
      <td>${lead.email} </td>
      <td>${lead.phone}</td>
      <td>${btnRemove}</td>
    </tr>`;
    document.getElementById('tab-leads').innerHTML = li;
  });
};

const showForm = (formId, firstInputId) => {
  document.getElementById(formId).style.display = 'block';
  if (firstInputId) {
    document.getElementById(firstInputId).focus();
    validateFormLead();
  }
};

const isValidEmail = (email) => {
  const re =
    /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  return re.test(String(email).toLowerCase());
};

const setError = (element, message) => {
  const inputControl = element.parentElement;
  const errorDisplay = inputControl.querySelector('.error');

  errorDisplay.innerText = message;
  inputControl.classList.add('error');
  inputControl.classList.remove('success');
};

const setSuccess = (element) => {
  const inputControl = element.parentElement;
  const errorDisplay = inputControl.querySelector('.error');

  errorDisplay.innerText = '';
  inputControl.classList.add('success');
  inputControl.classList.remove('error');
};

const validateFormLead = () => {
  const nameValue = document.getElementById('fname').value.trim();
  const emailValue = document.getElementById('femail').value.trim();
  const phoneValue = document.getElementById('fphone').value.trim();

  setError(document.getElementById('fname'), 'Username is required');
};
