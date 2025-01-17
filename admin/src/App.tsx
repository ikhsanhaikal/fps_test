import {
  Admin,
  Create,
  Datagrid,
  DataProvider,
  DateField,
  DateInput,
  DeleteButton,
  fetchUtils,
  List,
  NumberInput,
  ReferenceArrayField,
  ReferenceField,
  ReferenceInput,
  ReferenceOneField,
  required,
  Resource,
  SimpleForm,
  TextField,
  TextInput,
  WithRecord,
} from "react-admin";

import queryString from "query-string";
import { useState } from "react";

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
        <ReferenceInput source="kategori_id" reference={"kategori"} />
        {/* <ReferenceInput source="status_id" reference={"status"} /> */}
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
        {/* <DeleteButton /> */}
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
    console.log("url: ", url);
    const resp = await httpClient(url, { signal: params.signal });
    console.log(resp);

    return { data: resp.json.data };
  },
};

export default App;
