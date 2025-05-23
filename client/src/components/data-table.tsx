import {
	ColumnDef,
	getCoreRowModel,
	useReactTable,
	flexRender,
} from "@tanstack/react-table";
import type { PropsWithChildren } from 'react';
import {
	Table,
	TableBody,
	TableCell,
	TableHead,
	TableHeader,
	TableRow,
} from "@/components/ui/table";

type DataTableProps<TData, TValue> = {
	columns: ColumnDef<TData, TValue>[];
	data: TData[];
} & PropsWithChildren;

export function DataTable<TData, TValue>({
	columns,
	data,
	children,
}: DataTableProps<TData, TValue>) {

	const table = useReactTable({
		data,
		columns,
		debugTable: true,
		getCoreRowModel: getCoreRowModel()
	});

	return (
		<div className="container mx-auto">
			<div className='flex w-full items-end py-2 gap-1'>
				{children}
			</div>
			<div className="rounded-md border overflow-auto">
				<Table>
					<TableHeader className='bg-sidebar'>
						{table.getHeaderGroups().map((headerGroup) => (
							<TableRow key={headerGroup.id}>
								{headerGroup.headers.map((header) => {
									return (
										<TableHead key={header.id}>
											{header.isPlaceholder
												? null
												: flexRender(
													header.column.columnDef.header,
													header.getContext()
												)}
										</TableHead>
									);
								})}
							</TableRow>
						))}
					</TableHeader>
					<TableBody>
						{table.getRowModel().rows?.length ? (
							table.getRowModel().rows.map((row) => (
								<TableRow
									key={row.id}
									data-state={row.getIsSelected() && "selected"}
								>
									{row.getVisibleCells().map((cell) => (
										<TableCell key={cell.id}>
											{flexRender(cell.column.columnDef.cell, cell.getContext())}
										</TableCell>
									))}
								</TableRow>
							))
						) : (
							<TableRow>
								<TableCell colSpan={columns.length} className="h-24 text-center">
									No results.
								</TableCell>
							</TableRow>
						)}
					</TableBody>
				</Table>
			</div>
		</div>
	);
};