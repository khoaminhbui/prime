export class Response {
    errorCode: number;
    message: string;
    value: number;

    constructor(errorCode: number, message: string, value: number) {
        this.errorCode = errorCode;
        this.message = message;
        this.value = value;
    }
}