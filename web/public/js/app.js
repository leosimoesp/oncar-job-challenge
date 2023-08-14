'use strict';

const serverUrl = 'http://localhost:8080/api';

const listVehicles = async () => {
  await axios
    .post(`${serverUrl}/vehicles`)
    .then(function (response) {
      const vehicles = response.data;
      let li = `<tr><th>Marca</th><th>Modelo</th><th>Ano</th><th>Preço</th><th>Ação</th></tr>`;
      vehicles.forEach((vehicle) => {
        const data = btoa(JSON.stringify(vehicle));
        const button = `<button onclick="listLeads('${data}');">leads</button>`;
        li += `<tr>
          <td>${vehicle.brand}</td>
          <td>${vehicle.model} </td>
          <td>${vehicle.year}</td>
          <td>${vehicle.fmtPrice}</td>
          <td>${button}</td>
        </tr>`;
      });
      document.getElementById('tab-vehicles').innerHTML = li;
    })
    .catch(function (error) {
      console.log(error);
    });
};

const saveLead = async (e) => {
  e.preventDefault();
  const form = Object.fromEntries(new FormData(e.target).entries());

  const lead = {};
  lead.vehicleId = form.leadVehicleId;
  lead.email = form.email;
  lead.name = form.name;
  lead.phone = form.phone;

  const submitSuccess = document.querySelector('.submit-success');
  const submitError = document.querySelector('.submit-error');

  await axios
    .post(`${serverUrl}/leads`, lead)
    .then(function (response) {
      const status = response.status;

      if (status === 201) {
        submitError.style.display = 'none';
        submitSuccess.style.display = 'block';
        getVehicleLeads(lead.vehicleId);
        clearForm('flead');
      }
    })
    .catch(function (error) {
      submitSuccess.style.display = 'none';
      console.log(error);
      submitError.style.display = 'block';
      if (error.response.data.message) {
        submitError.innerHTML = `<p class="submit-error-text">${error.response.data.message}</p>`;
      } else {
        submitError.innerHTML = `<p class="submit-error-text">${JSON.stringify(
          error.response.data
        )}</p>`;
      }
    });
};

const getVehicleLeads = async (vehicleId) => {
  const leads = await axios
    .get(`${serverUrl}/leads/${vehicleId}`)
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
    });

  let li = `<tr><th>Nome</th><th>E-mail</th><th>Telefone</th><th>Ação</th></tr>`;
  leads.forEach((lead) => {
    const btnRemove = `<button onclick="alert('@TODO')">excluir</button>`;
    li += `<tr>
      <td>${lead.name}</td>
      <td>${lead.email} </td>
      <td>${lead.phone}</td>
      <td>${btnRemove}</td>
    </tr>`;
  });
  document.getElementById('tab-leads').innerHTML = li;
};

const listLeads = async (vehicle) => {
  const parsed = JSON.parse(atob(vehicle));
  let div = `
    <div>Marca: ${parsed.brand}</div>
    <div>Model: ${parsed.model}</div>
    <div>Ano: ${parsed.year}</div>
    <div>Preço: ${parsed.fmtPrice}</div>
    <br/>
    <div><center><h2>Leads</h2></center></div>
    <div class="button-container">
      <div class="button-inner"><button onclick="showForm('lead-container','fname');">novo</button></div>
      <div class="button-inner"><button onclick="navigate('list-vehicles')">voltar</button></div>
    </div>
  `;
  document.getElementById('tab-vehicles').style.display = 'none';
  document.getElementById('vehicle-leads').innerHTML = div;

  await getVehicleLeads(parsed.id);

  document.getElementById('fvehicleid').value = parsed.id;
  const form = document.getElementById('flead');
  form.addEventListener('submit', saveLead);
};

const clearForm = (form) => {
  document.getElementById(form).reset();
};

const showForm = (formId, firstInputId) => {
  clearForm('flead');
  document.getElementById(formId).style.display = 'block';
  if (firstInputId) {
    document.getElementById(firstInputId).focus();
  }
};

const navigate = (htmlName) => {
  location.href = htmlName;
};

const isValidEmail = (email) => {
  const re =
    /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  return re.test(String(email).toLowerCase());
};

const isValidPhone = (phone) => {
  const value = phone.replaceAll(' ', '');
  const onlyNumbers = new RegExp('^[0-9]+$');
  const validOnlyNums = onlyNumbers.test(value);
  const size = value.length;
  return size <= 11 && size > 9 && validOnlyNums;
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
  const elemName = document.getElementById('fname');
  const nameValue = elemName.value.trim();
  const elemEmail = document.getElementById('femail');
  const emailValue = elemEmail.value.trim();
  const elemPhone = document.getElementById('fphone');
  const phoneValue = elemPhone.value.trim();
  let hasErrors = document.getElementById('fsubmit').disabled;

  if (nameValue === '') {
    setError(elemName, 'Nome é obrigatório');
    hasErrors = true;
  } else {
    setSuccess(elemName);
    hasErrors = false;
  }

  if (emailValue === '') {
    setError(elemEmail, 'E-mail é obrigatório');
    hasErrors = true;
  } else {
    setSuccess(elemEmail);
    hasErrors = false;
  }

  if (emailValue !== '' && !isValidEmail(emailValue)) {
    setError(elemEmail, 'E-mail é inválido');
    hasErrors = true;
  } else if (emailValue !== '') {
    setSuccess(elemEmail);
    hasErrors = false;
  }

  if (phoneValue === '') {
    setError(elemPhone, 'Telefone é obrigatório');
    hasErrors = true;
  } else {
    setSuccess(elemPhone);
    hasErrors = false;
  }

  if (phoneValue !== '' && !isValidPhone(phoneValue)) {
    setError(elemPhone, 'Telefone é inválido');
    hasErrors = true;
  } else if (phoneValue !== '') {
    setSuccess(elemPhone);
    hasErrors = false;
  }

  if (!hasErrors) {
    document.getElementById('fsubmit').disabled = false;
  } else {
    document.getElementById('fsubmit').disabled = true;
  }
};

const closeElem = (elemId) => {
  document.getElementById(elemId).style.display = 'none';
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
