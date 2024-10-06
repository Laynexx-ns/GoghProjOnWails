import {useEffect, useState} from "react";
import {GetPublicRepositories} from "../../../wailsjs/go/main/App.js";
import RepositoryDetails from "./RepositoryDetails.jsx";
import {message} from "antd";
import MasterDetail from "../MasterDetail.jsx";

const PublicRepositories = () =>{
    const [repositories, setRepositories] = useState([]);
    const [messageApi, contextHolder] = message.useMessage();


    //getting repositories effect
    useEffect = (() => {
        const getRepositories = async () => {
            GetPublicRepositories()
                .then( //if it's ok
                (repositories) => {
                    setRepositories(repositories)
                })
                .catch( //if it's not ok((((
                (error) => {
                messageApi.open({ //say fuck off to user
                    type: error, content: error
                });
            });
        }
        getRepositories();
    }, []);//end of useEffect


    const title = "Public Repositories";
    const getItemDescription = (repository) => repository.description;
    const detailLayout = (repository) => <RepositoryDetails repository={repository}/> ;


    //return master detail
    return(
        <>
            {contextHolder}
            <MasterDetail
                title={title}
                items={repositories}
                getItemDescription={getItemDescription()}
                detailLayout={detailLayout()}
            />
        </>
    )
};

export default PublicRepositories;
