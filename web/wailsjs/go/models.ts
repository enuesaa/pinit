export namespace usecase {
	
	export class ChatRequest {
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new ChatRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	    }
	}
	export class ChatResponse {
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new ChatResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	    }
	}
	export class CreateBinderRequest {
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateBinderRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	    }
	}
	export class CreateNoteRequest {
	    binderId: number;
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateNoteRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.binderId = source["binderId"];
	        this.content = source["content"];
	    }
	}
	export class ListActionsItem {
	    name: string;
	    template: string;
	
	    static createFrom(source: any = {}) {
	        return new ListActionsItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.template = source["template"];
	    }
	}
	export class ListBindersItem {
	    id: string;
	    name: string;
	    category: string;
	    archivedAt: string;
	    createdAt: string;
	    updatedAt: string;
	
	    static createFrom(source: any = {}) {
	        return new ListBindersItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.category = source["category"];
	        this.archivedAt = source["archivedAt"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	    }
	}
	export class ListNotesItem {
	    id: string;
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new ListNotesItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.content = source["content"];
	    }
	}
	export class ServeCreateResponse {
	    id: number;
	
	    static createFrom(source: any = {}) {
	        return new ServeCreateResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	    }
	}

}

