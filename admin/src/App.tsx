import {
  Admin,
  AutocompleteInput,
  Create,
  Datagrid,
  DataProvider,
  DateField,
  DeleteButton,
  fetchUtils,
  List,
  NumberInput,
  ReferenceField,
  ReferenceInput,
  required,
  Resource,
  SimpleForm,
  TextField,
  TextInput,
} from "react-admin";

import queryString from "query-string";

function App() {
  return (
    <Admin dataProvider={dataProvider}>
      <Resource
        name="produk"
        list={ListProduk}
        create={CreateProduk}
        hasShow={true}
      />
    </Admin>
  );
}

function CreateProduk() {
  return (
    <Create>
      <SimpleForm>
        <TextInput source="nama" validate={[required()]} />
        <ReferenceInput source="kategori_id" reference="kategori">
          <AutocompleteInput optionText="nama_kategori" />
        </ReferenceInput>
        <ReferenceInput source="status_id" reference="status">
          <AutocompleteInput optionText="nama_status" />
        </ReferenceInput>
        <NumberInput
          source="harga"
          label="Harga produk"
          validate={[required()]}
        />
      </SimpleForm>
    </Create>
  );
}

function ListProduk() {
  return (
    <List exporter={false}>
      <Datagrid>
        <TextField source="id" />
        <TextField source="nama_produk" />
        <TextField source="harga" />
        <ReferenceField
          source="kategori_id"
          reference="kategori"
          label="kategori"
        >
          <TextField source="nama_kategori" />
        </ReferenceField>
        <ReferenceField source="status_id" reference="status" label="status">
          <TextField source="nama_status" />
        </ReferenceField>
        <DateField source="created_at" />
        <DeleteButton mutationMode="pessimistic" />
      </Datagrid>
    </List>
  );
}

const apiUrl = "http://localhost:3000";
const httpClient = fetchUtils.fetchJson;

const dataProvider: DataProvider = {
  getList: async (resource, params) => {
    let { page, perPage }: { page: number; perPage: number } =
      params.pagination;
    const url = `${apiUrl}/${resource}?${queryString.stringify({
      page,
      perPage,
    })}`;
    const resp = await httpClient(url, { signal: params.signal });
    console.log("resp: ", resp.json.data);
    return {
      data: resp.json.data,
      total: resp.json.total,
    };
  },
  getMany: async (resource, params) => {
    const query = {
      ids: params.ids,
    };
    const url = `${apiUrl}/${resource}?${queryString.stringify(query)}`;
    // console.log("url: ", url);
    const resp = await httpClient(url, { signal: params.signal });
    // console.log(resp);

    return { data: resp.json.data };
  },
  create: async (resource, params) => {
    const response = await fetch(`${apiUrl}/${resource}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json;charset=utf-8",
      },
      body: JSON.stringify(params.data),
    });

    const result = await response.json();
    // console.log("result: ", result);
    return { data: result.data };
  },
  delete: async (resource, params) => {
    const url = `${apiUrl}/${resource}/${params.id}`;
    const { json } = await httpClient(url, {
      method: "DELETE",
    });
    return { data: json.data };
  },
};

export default App;
