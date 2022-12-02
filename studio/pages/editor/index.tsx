import { useEffect, useRef } from 'react';
import { debounce } from 'lodash';
import 'cherry-markdown/dist/cherry-markdown.css'

const FormData = require('form-data');

const Index = () => {

    const cherryInstance = useRef<any>();
    const formData = new FormData();
    formData.append('fileID', 'intro.md');
    formData.append('path', 'docusaurus/docs');

    const handleEditorChange = debounce((value: any, ...args: any[]) => {
        formData.set('content', value);
        fetch(`/favorite/catalogue/writeFile`, {
            method: 'POST',
            body: formData
        }).then((res) => res.json())
        .then((data) => {
            console.log('handleEditorChange:', data);
        });
    }, 500)

    useEffect(() => {
        if (!cherryInstance.current) {
            const Cherry = require('cherry-markdown').default;

            fetch(`/favorite/catalogue/getFile`, {
                method: 'POST',
                body: formData
            }).then((res) => res.json())
            .then((data) => {
                cherryInstance.current = new Cherry({
                    id: 'markdown-container',
                    value: data.data,
                    callback: {
                        afterChange: handleEditorChange
                    }
                });
            });
        }

    }, [])

    return (
        <div>
            <div>
                <div id="markdown-container"></div>
            </div>
        </div>
    )
}

export default Index;